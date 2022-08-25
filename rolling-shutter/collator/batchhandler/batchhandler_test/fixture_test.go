package batchhandler_test

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	gocmp "github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	txtypes "github.com/shutter-network/txtypes/types"
	"gotest.tools/assert"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/collator/batchhandler"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/collator/batchhandler/sequencer"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/collator/cltrdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/collator/config"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/epochid"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/testdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/shmsg"
)

// From the gocmp Example on Reporter interface:

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  gocmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps gocmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs gocmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

func assertEqual(t *testing.T, cancel func(error), x, y any, opts ...gocmp.Option) {
	t.Helper()
	rep := &DiffReporter{}
	opts = append(opts, gocmp.Reporter(rep))
	res := gocmp.Equal(x, y, opts...)
	if !res {
		// TODO this does not have the AST based reporting
		// of the expression being evaluated.
		// This would be nice to have, but is not a must.
		msg := "assertEqual failed " + rep.String()
		t.Error(msg)
		cancel(errors.New(msg))
	}
}

func compareBigInt(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func compareByte(a, b []byte) bool {
	return bytes.Equal(a, b)
}

func newTestConfig(t *testing.T) config.Config {
	t.Helper()

	ethereumKey, err := ethcrypto.GenerateKey()
	assert.NilError(t, err)
	return config.Config{
		EthereumURL:         "http://127.0.0.1:8454",
		SequencerURL:        "http://127.0.0.1:8455",
		EthereumKey:         ethereumKey,
		ExecutionBlockDelay: uint32(5),
		InstanceID:          123,
		EpochDuration:       2 * time.Second,
	}
}

type TestParams struct {
	GasLimit       uint64
	InitialBalance *big.Int
	BaseFee        *big.Int
	TxGasTipCap    *big.Int
	TxGasFeeCap    *big.Int
	InitialEpochID epochid.EpochID
	EpochDuration  time.Duration
}

type Fixture struct {
	Cfg          config.Config
	EthL1Server  *sequencer.MockEthServer
	EthL2Server  *sequencer.MockEthServer
	BatchHandler *batchhandler.BatchHandler
	MakeTx       func(batchIndex, nonce, gas int) ([]byte, []byte)
	Address      common.Address
	Coinbase     common.Address
}

func Setup(ctx context.Context, t *testing.T, params TestParams) *Fixture {
	t.Helper()

	cfg := newTestConfig(t)
	if params.EpochDuration != 0 {
		cfg.EpochDuration = params.EpochDuration
	}
	ethL1 := sequencer.RunMockEthServer(t)
	ethL2 := sequencer.RunMockEthServer(t)

	cfg.EthereumURL = ethL1.URL
	cfg.SequencerURL = ethL2.URL

	db, dbpool, dbteardown := testdb.NewCollatorTestDB(ctx, t)
	t.Cleanup(ethL1.Teardown)
	t.Cleanup(ethL2.Teardown)
	t.Cleanup(dbteardown)

	address := ethcrypto.PubkeyToAddress(cfg.EthereumKey.PublicKey)
	chainID := big.NewInt(0)
	gasLimit := params.GasLimit
	signer := txtypes.LatestSignerForChainID(chainID)
	coinbase := common.HexToAddress("0x0000000000000000000000000000000000000000")

	// Set the values on the dummy rpc server
	ethL1.SetBlockNumber(1)
	ethL2.SetBalance(address, params.InitialBalance, "latest")
	ethL2.SetBalance(coinbase, big.NewInt(0), "latest")
	ethL2.SetNonce(address, uint64(0), "latest")
	ethL2.SetChainID(chainID)
	ethL2.SetBlock(params.BaseFee, gasLimit, "latest")

	// set initial ("next") epoch id and block number manually,
	// this is usually done in the collator and not in the handler

	err := db.SetNextBatch(ctx, cltrdb.SetNextBatchParams{
		EpochID:       params.InitialEpochID.Bytes(),
		L1BlockNumber: 42,
	})
	assert.NilError(t, err)

	// New batch handler, this will already query the eth-server
	bh, err := batchhandler.NewBatchHandler(cfg, dbpool)
	assert.NilError(t, err)

	makeTx := func(batchIndex, nonce, gas int) ([]byte, []byte) {
		// construct a valid transaction
		txData := &txtypes.ShutterTx{
			ChainID:          chainID,
			Nonce:            uint64(nonce),
			GasTipCap:        params.TxGasTipCap,
			GasFeeCap:        params.TxGasFeeCap,
			Gas:              uint64(gas),
			EncryptedPayload: []byte("foo"),
			BatchIndex:       uint64(batchIndex),
		}
		tx, err := txtypes.SignNewTx(cfg.EthereumKey, signer, txData)
		assert.NilError(t, err)

		// marshal tx to bytes
		txBytes, err := tx.MarshalBinary()
		assert.NilError(t, err)
		return txBytes, tx.Hash().Bytes()
	}

	return &Fixture{
		Cfg:          cfg,
		EthL2Server:  ethL2,
		EthL1Server:  ethL1,
		BatchHandler: bh,
		MakeTx:       makeTx,
		Address:      address,
		Coinbase:     coinbase,
	}
}

// `p2pMessagingMock` shortcuts the entire P2P messaging layer that would
// normally handle the communication with the keyper-set.
// If the `p2pMessagingMock` receives a DecryptionTrigger,
// it will wait a fixed amount of time and then
// call the BatchHandler's HandleDecryptionKey with a fixed decryption-key
// for that epoch.
// The data of the DecryptionTrigger is not fully validated.
func p2pMessagingMock(
	t *testing.T,
	ctx context.Context,
	failFunc func(error),
	config config.Config,
	batchHandler *batchhandler.BatchHandler,
) error {
	t.Helper()

	messages := batchHandler.Messages()
	for {
		select {
		case out, ok := <-messages:
			if !ok {
				log.Debug().Msg("P2P message mock: messages closed from outside")
				return nil
			}
			// Emulate communication overhead etc.
			time.Sleep(500 * time.Millisecond)
			typ := out.ProtoReflect().Type().Descriptor().FullName()
			if typ == "shmsg.DecryptionTrigger" {
				trigger, _ := out.(*shmsg.DecryptionTrigger)
				epoch, _ := epochid.BytesToEpochID(trigger.EpochID)

				assertEqual(t, failFunc, trigger.InstanceID, config.InstanceID)
				address := ethcrypto.PubkeyToAddress(config.EthereumKey.PublicKey)
				signatureCorrect, err := shmsg.VerifySignature(trigger, address)
				assertEqual(t, failFunc, err, nil)
				assertEqual(t, failFunc, signatureCorrect, true)

				// collator successfully received the decryption-key via P2P messaging,
				// pass to the batch-handler
				err = batchHandler.HandleDecryptionKey(ctx, epoch, []byte("decrkey"))
				assertEqual(t, failFunc, err, nil)
			}

		case <-ctx.Done():
			err := ctx.Err()
			if err == context.Canceled {
				return nil
			}
			return err
		}
	}
}