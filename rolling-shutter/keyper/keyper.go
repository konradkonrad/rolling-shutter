// Package keyper contains the keyper implementation
package keyper

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/sync/errgroup"

	"github.com/shutter-network/shutter/shuttermint/keyper/fx"
	"github.com/shutter-network/shutter/shuttermint/keyper/kprdb"
	"github.com/shutter-network/shutter/shuttermint/p2p"
)

var GossipTopicNames = []string{"decryptionTrigger", "decryptionKey"}

type keyper struct {
	config            Config
	dbpool            *pgxpool.Pool
	shuttermintClient client.Client
	messageSender     fx.RPCMessageSender
	shuttermintState  *ShuttermintState
	p2p               *p2p.P2P
}

// linkConfigToDB ensures that we use a database compatible with the given config. On first use
// it stores the config's ethereum address into the database. On subsequent uses it compares the
// stored value and raises an error if it doesn't match.
func linkConfigToDB(ctx context.Context, config Config, dbpool *pgxpool.Pool) error {
	const addressKey = "ethereum address"
	cfgAddress := config.Address().Hex()
	queries := kprdb.New(dbpool)
	dbAddr, err := queries.GetMeta(ctx, addressKey)
	if err == pgx.ErrNoRows {
		return queries.InsertMeta(ctx, kprdb.InsertMetaParams{
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

func Run(ctx context.Context, config Config) error {
	dbpool, err := pgxpool.Connect(ctx, config.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}
	defer dbpool.Close()

	err = kprdb.ValidateKeyperDB(ctx, dbpool)
	if err != nil {
		return err
	}
	err = linkConfigToDB(ctx, config, dbpool)
	if err != nil {
		return err
	}
	shuttermintClient, err := http.New(config.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}
	messageSender := fx.NewRPCMessageSender(shuttermintClient, config.SigningKey)

	k := keyper{
		config:            config,
		dbpool:            dbpool,
		shuttermintClient: shuttermintClient,
		messageSender:     messageSender,
		shuttermintState:  NewShuttermintState(config),
		p2p: p2p.New(p2p.Config{
			ListenAddr:     config.ListenAddress,
			PeerMultiaddrs: config.PeerMultiaddrs,
			PrivKey:        config.P2PKey,
		}),
	}
	return k.run(ctx)
}

func (kpr *keyper) run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return kpr.p2p.Run(ctx, GossipTopicNames)
	})
	group.Go(func() error {
		return kpr.operateShuttermint(ctx)
	})
	return group.Wait()
}

func (kpr *keyper) operateShuttermint(ctx context.Context) error {
	for {
		err := SyncAppWithDB(ctx, kpr.shuttermintClient, kpr.dbpool, kpr.shuttermintState)
		if err != nil {
			return err
		}
		err = SendShutterMessages(ctx, kprdb.New(kpr.dbpool), &kpr.messageSender)
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
