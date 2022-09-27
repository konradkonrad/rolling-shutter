// Package l2client provides some convenience functions to interact with the sequencer.
package l2client

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	txtypes "github.com/shutter-network/txtypes/types"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/epochid"
)

// GetBatchIndex retrieves the current batch index from the sequencer.
func GetBatchIndex(ctx context.Context, l2Client *rpc.Client) (epochid.EpochID, error) {
	var epochID epochid.EpochID

	f := func() (*string, error) {
		var result string
		log.Debug().Msg("polling batch-index from sequencer")
		err := l2Client.CallContext(ctx, &result, "shutter_getBatchIndex")
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	result, err := medley.Retry(ctx, f)
	if err != nil {
		return epochID, errors.Wrapf(err, "can't retrieve batch-index from sequencer")
	}

	epochID, err = epochid.HexToEpochID(*result)
	if err != nil {
		return epochID, errors.Wrap(err, "can't decode batch-index")
	}
	return epochID, nil
}

// SendTransaction sends a transaction to the sequencer. It uses the raw rpc.Client instead of the
// usual ethclient.Client wrapper because we want to use the modified txtypes marshaling here
// instead of the one from the go-ethereum repository.
func SendTransaction(ctx context.Context, client *rpc.Client, tx *txtypes.Transaction) error {
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	f := func() (string, error) {
		var result string
		//
		err := client.CallContext(ctx, &result, "eth_sendRawTransaction", hexutil.Encode(data))
		if err != nil {
			return result, err
		}
		return result, nil
	}
	_, err = medley.Retry(ctx, f)
	if err != nil {
		return errors.Wrap(err, "can't send transaction")
	}
	return err
}