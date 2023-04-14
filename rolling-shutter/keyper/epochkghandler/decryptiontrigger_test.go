package epochkghandler

import (
	"bytes"
	"context"
	"testing"

	"gotest.tools/assert"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/db/kprdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/epochid"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/testdb"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2p"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2pmsg"
)

func TestHandleDecryptionTriggerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	db, dbpool, closedb := testdb.NewKeyperTestDB(ctx, t)
	defer closedb()

	epochID := epochid.Uint64ToEpochID(50)
	keyperIndex := uint64(1)

	initializeEon(ctx, t, db, keyperIndex)
	var handler p2p.MessageHandler = &DecryptionTriggerHandler{config: config, dbpool: dbpool}
	// send decryption key share when first trigger is received
	trigger := &p2pmsg.DecryptionTrigger{
		EpochID:    epochID.Bytes(),
		InstanceID: 0,
	}
	msgs, err := handler.HandleMessage(ctx, trigger)
	assert.NilError(t, err)
	share, err := db.GetDecryptionKeyShare(ctx, kprdb.GetDecryptionKeyShareParams{
		EpochID:     epochID.Bytes(),
		KeyperIndex: int64(keyperIndex),
	})
	assert.NilError(t, err)
	assert.Check(t, len(msgs) == 1)
	msg, ok := msgs[0].(*p2pmsg.DecryptionKeyShare)
	assert.Check(t, ok)
	assert.Check(t, msg.InstanceID == config.GetInstanceID())
	assert.Check(t, bytes.Equal(msg.EpochID, epochID.Bytes()))
	assert.Check(t, msg.KeyperIndex == keyperIndex)
	assert.Check(t, bytes.Equal(msg.Share, share.DecryptionKeyShare))

	// don't send share when trigger is received again
	msgs, err = handler.HandleMessage(ctx, trigger)
	assert.NilError(t, err)
	assert.Check(t, len(msgs) == 0)
}
