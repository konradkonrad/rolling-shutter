// Package snapshot_keyper contains the snapshot specific keyper implementation
package snapshot_keyper

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/tendermint/tendermint/rpc/client"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/chainobserver"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/contract/deployment"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/db/chainobsdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/db/kprdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/db/metadb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper/epochkghandler"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper/fx"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper/kprapi"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper/smobserver"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/eventsyncer"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/retry"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/service"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2p"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2pmsg"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/shdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/shmsg"
)

type snapshot_keyper struct {
	config            *Config
	dbpool            *pgxpool.Pool
	shuttermintClient client.Client
	messageSender     fx.RPCMessageSender
	l1Client          *ethclient.Client
	contracts         *deployment.Contracts

	shuttermintState *smobserver.ShuttermintState
	p2p              *p2p.P2PHandler
}

func New(config Config) service.Service {
	return &snapshot_keyper{config: &config}
}

// linkConfigToDB ensures that we use a database compatible with the given config. On first use
// it stores the config's ethereum address into the database. On subsequent uses it compares the
// stored value and raises an error if it doesn't match.
func linkConfigToDB(ctx context.Context, config *Config, dbpool *pgxpool.Pool) error {
	const addressKey = "ethereum address"
	cfgAddress := config.GetAddress().Hex()
	queries := metadb.New(dbpool)
	dbAddr, err := queries.GetMeta(ctx, addressKey)
	if err == pgx.ErrNoRows {
		return queries.InsertMeta(ctx, metadb.InsertMetaParams{
			Key:   addressKey,
			Value: cfgAddress,
		})
	} else if err != nil {
		return err
	}

	if dbAddr != cfgAddress {
		return errors.Errorf(
			"database linked to wrong address %s, config address is %s",
			dbAddr, cfgAddress)
	}
	return nil
}

func (kpr *snapshot_keyper) Start(ctx context.Context, runner service.Runner) error {
	config := kpr.config
	dbpool, err := pgxpool.Connect(ctx, config.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}
	runner.Defer(dbpool.Close)
	shdb.AddConnectionInfo(log.Info(), dbpool).Msg("connected to database")

	l1Client, err := ethclient.Dial(config.EthereumURL)
	if err != nil {
		return err
	}
	/*l2Client, err := ethclient.Dial(config.ContractsURL)
	if err != nil {
		return err
	}*/
	contracts, err := deployment.NewContracts(l1Client, config.DeploymentDir)
	if err != nil {
		return err
	}

	err = kprdb.ValidateKeyperDB(ctx, dbpool)
	if err != nil {
		return err
	}
	err = linkConfigToDB(ctx, config, dbpool)
	if err != nil {
		return err
	}
	shuttermintClient, err := tmhttp.New(config.ShuttermintURL)
	if err != nil {
		return err
	}
	messageSender := fx.NewRPCMessageSender(shuttermintClient, config.SigningKey)

	p2pHandler := p2p.New(p2p.Config{
		ListenAddrs:    config.ListenAddresses,
		BootstrapPeers: config.CustomBootstrapAddresses,
		PrivKey:        config.P2PKey,
		Environment:    p2p.Production,
	})

	kpr.dbpool = dbpool
	kpr.shuttermintClient = shuttermintClient
	kpr.messageSender = messageSender
	kpr.l1Client = l1Client
	kpr.contracts = contracts
	kpr.shuttermintState = smobserver.NewShuttermintState(config)
	kpr.p2p = p2pHandler

	kpr.setupP2PHandler()
	return runner.StartService(kpr.getServices()...)
}

func (snkpr *snapshot_keyper) setupP2PHandler() {
	snkpr.p2p.AddMessageHandler(
		epochkghandler.NewDecryptionKeyHandler(snkpr.config, snkpr.dbpool),
		epochkghandler.NewDecryptionKeyShareHandler(snkpr.config, snkpr.dbpool),
		epochkghandler.NewDecryptionTriggerHandler(snkpr.config, snkpr.dbpool),
		epochkghandler.NewEonPublicKeyHandler(snkpr.config, snkpr.dbpool),
	)
}

func (snkpr *snapshot_keyper) getServices() []service.Service {
	services := []service.Service{
		snkpr.p2p,
		service.ServiceFn{Fn: snkpr.operateShuttermint},
		service.ServiceFn{Fn: snkpr.broadcastEonPublicKeys},
		service.ServiceFn{Fn: snkpr.handleContractEvents},
	}

	if snkpr.config.HTTPEnabled {
		services = append(services, kprapi.NewHTTPService(snkpr.dbpool, snkpr.config, snkpr.p2p))
	}
	return services
}

func (kpr *snapshot_keyper) handleContractEvents(ctx context.Context) error {
	events := []*eventsyncer.EventType{
		kpr.contracts.KeypersConfigsListNewConfig,
		kpr.contracts.CollatorConfigsListNewConfig,
	}
	return chainobserver.New(kpr.contracts, kpr.dbpool).Observe(ctx, events)
}

func (snkpr *snapshot_keyper) handleOnChainChanges(ctx context.Context, tx pgx.Tx, l1BlockNumber uint64) error {
	err := snkpr.handleOnChainKeyperSetChanges(ctx, tx)
	if err != nil {
		return err
	}
	err = snkpr.sendNewBlockSeen(ctx, tx, l1BlockNumber)
	if err != nil {
		return err
	}
	return nil
}

// sendNewBlockSeen sends shmsg.NewBlockSeen messages to the shuttermint chain. This function sends
// NewBlockSeen messages to the shuttermint chain, so that the chain can start new batch configs if
// enough keypers have seen a block past the start block of some BatchConfig. We only send messages
// when the current block we see, could lead to a batch config being started.
func (snkpr *snapshot_keyper) sendNewBlockSeen(ctx context.Context, tx pgx.Tx, l1BlockNumber uint64) error {
	q := kprdb.New(tx)
	lastBlock, err := q.GetLastBlockSeen(ctx)
	if err != nil {
		return err
	}

	count, err := q.CountBatchConfigsInBlockRange(ctx,
		kprdb.CountBatchConfigsInBlockRangeParams{
			StartBlock: lastBlock,
			EndBlock:   int64(l1BlockNumber),
		})
	if err != nil {
		return err
	}
	if count == 0 {
		return nil
	}

	blockSeenMsg := shmsg.NewBlockSeen(l1BlockNumber)
	err = q.ScheduleShutterMessage(ctx, "block seen", blockSeenMsg)
	if err != nil {
		return err
	}
	err = q.SetLastBlockSeen(ctx, int64(l1BlockNumber))
	if err != nil {
		return err
	}
	log.Info().Uint64("block-number", l1BlockNumber).Msg("block seen")
	return nil
}

// handleOnChainKeyperSetChanges looks for changes in the keyper_set table.
func (snkpr *snapshot_keyper) handleOnChainKeyperSetChanges(ctx context.Context, tx pgx.Tx) error {
	q := kprdb.New(tx)
	latestBatchConfig, err := q.GetLatestBatchConfig(ctx)
	if err == pgx.ErrNoRows {
		log.Print("no batch config found in tendermint")
		return nil
	} else if err != nil {
		return err
	}

	cq := chainobsdb.New(tx)
	keyperSet, err := cq.GetKeyperSetByKeyperConfigIndex(ctx, int64(latestBatchConfig.KeyperConfigIndex)+1)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		return err
	}

	lastSent, err := q.GetLastBatchConfigSent(ctx)
	if err != nil {
		return err
	}
	if lastSent == keyperSet.KeyperConfigIndex {
		return nil
	}

	lastBlock, err := q.GetLastBlockSeen(ctx)
	if err != nil {
		return err
	}

	if keyperSet.ActivationBlockNumber-lastBlock > snkpr.config.DKGStartBlockDelta {
		log.Info().Interface("keyper-set", keyperSet).Msg("not yet submitting config")
		return nil
	}

	err = q.SetLastBatchConfigSent(ctx, keyperSet.KeyperConfigIndex)
	if err != nil {
		return nil
	}

	keypers, err := shdb.DecodeAddresses(keyperSet.Keypers)
	if err != nil {
		return err
	}
	log.Info().Interface("keyper-set", keyperSet).Msg("have a new config to be scheduled")
	batchConfigMsg := shmsg.NewBatchConfig(
		uint64(keyperSet.ActivationBlockNumber),
		keypers,
		uint64(keyperSet.Threshold),
		uint64(keyperSet.KeyperConfigIndex),
	)
	err = q.ScheduleShutterMessage(ctx, "new batch config", batchConfigMsg)
	if err != nil {
		return err
	}
	return nil
}

func (snkpr *snapshot_keyper) operateShuttermint(ctx context.Context) error {
	for {
		l1BlockNumber, err := retry.FunctionCall(ctx, snkpr.l1Client.BlockNumber)
		if err != nil {
			return err
		}

		err = smobserver.SyncAppWithDB(ctx, snkpr.shuttermintClient, snkpr.dbpool, snkpr.shuttermintState)
		if err != nil {
			return err
		}
		err = snkpr.dbpool.BeginFunc(ctx, func(tx pgx.Tx) error {
			return snkpr.handleOnChainChanges(ctx, tx, l1BlockNumber)
		})
		if err != nil {
			return err
		}

		err = fx.SendShutterMessages(ctx, kprdb.New(snkpr.dbpool), &snkpr.messageSender)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func (snkpr *snapshot_keyper) broadcastEonPublicKeys(ctx context.Context) error {
	for {
		eonPublicKeys, err := kprdb.New(snkpr.dbpool).GetAndDeleteEonPublicKeys(ctx)
		if err != nil {
			return err
		}
		for _, eonPublicKey := range eonPublicKeys {
			_, exists := kprdb.GetKeyperIndex(snkpr.config.GetAddress(), eonPublicKey.Keypers)
			if !exists {
				return errors.Errorf("own keyper index not found for Eon=%d", eonPublicKey.Eon)
			}
			msg, err := p2pmsg.NewSignedEonPublicKey(
				snkpr.config.InstanceID,
				eonPublicKey.EonPublicKey,
				uint64(eonPublicKey.ActivationBlockNumber),
				uint64(eonPublicKey.KeyperConfigIndex),
				uint64(eonPublicKey.Eon),
				snkpr.config.SigningKey,
			)
			if err != nil {
				return errors.Wrap(err, "error while signing EonPublicKey")
			}

			err = snkpr.p2p.SendMessage(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "error while broadcasting EonPublicKey")
			}
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}
