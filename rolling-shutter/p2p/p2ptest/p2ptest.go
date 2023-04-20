// p2ptest contains code for testing code implementing a p2p2.MessageHandler.
package p2ptest

import (
	"context"
	"testing"

	"gotest.tools/assert"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2p"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/p2pmsg"
)

func MustValidateMessageResult(
	t *testing.T,
	expectedResult bool,
	handler p2p.MessageHandler,
	ctx context.Context, //nolint:revive
	msg p2pmsg.Message,
) {
	t.Helper()
	ok, err := handler.ValidateMessage(ctx, msg)
	if expectedResult {
		assert.NilError(t, err, "validation returned error")
	}
	assert.Equal(t, ok, expectedResult,
		"validate failed ok=%t !=  msg=%+v type=%T", ok, expectedResult, msg, msg,
	)
}

func MustHandleMessage(
	t *testing.T,
	handler p2p.MessageHandler,
	ctx context.Context, //nolint:revive
	msg p2pmsg.Message,
) []p2pmsg.Message {
	t.Helper()
	MustValidateMessageResult(t, true, handler, ctx, msg)
	msgs, err := handler.HandleMessage(ctx, msg)
	assert.NilError(t, err)
	return msgs
}
