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

func TestHandleDecryptionKeyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.Background()
	db, dbpool, closedb := testdb.NewKeyperTestDB(ctx, t)
	defer closedb()

	eon := uint64(2)
	epochID := epochid.Uint64ToEpochID(50)
	keyperIndex := uint64(1)

	tkg := initializeEon(ctx, t, db, keyperIndex)

	var handler p2p.MessageHandler = &DecryptionKeyHandler{config: config, dbpool: dbpool}
	encodedDecryptionKey := tkg.EpochSecretKey(epochID).Marshal()

	// send a decryption key and check that it gets inserted
	msgs, err := handler.HandleMessage(ctx, &p2pmsg.DecryptionKey{
		InstanceID: 0,
		Eon:        eon,
		EpochID:    epochID.Bytes(),
		Key:        encodedDecryptionKey,
	})
	assert.NilError(t, err)
	assert.Check(t, len(msgs) == 0)
	key, err := db.GetDecryptionKey(ctx, kprdb.GetDecryptionKeyParams{
		Eon:     int64(eon),
		EpochID: epochID.Bytes(),
	})
	assert.NilError(t, err)
	assert.Check(t, bytes.Equal(key.DecryptionKey, encodedDecryptionKey))
}
