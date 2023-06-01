package snapshot

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shuttermint/p2p"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
	"github.com/shutter-network/shutter/shuttermint/snapshot/snptopics"
)

type (
	decryptionKey shmsg.DecryptionKey
	eonPublicKey  shmsg.EonPublicKey
	timedEpoch    shmsg.TimedEpoch
)

type message interface {
	implementsMessage()
	GetInstanceID() uint64
}

func (*decryptionKey) implementsMessage() {}
func (*eonPublicKey) implementsMessage()  {}
func (*timedEpoch) implementsMessage()    {}

func (d *decryptionKey) GetInstanceID() uint64 { return d.InstanceID }
func (e *eonPublicKey) GetInstanceID() uint64  { return e.InstanceID }
func (te *timedEpoch) GetInstanceID() uint64   { return te.InstanceID }

func unmarshalP2PMessage(msg *p2p.Message) (message, error) {
	if msg == nil {
		return nil, nil
	}
	switch msg.Topic {
	case snptopics.DecryptionKey:
		return unmarshalDecryptionKey(msg)
	case snptopics.EonPublicKey:
		return unmarshalEonPublicKey(msg)
	case snptopics.TimedEpoch:
		return unmarshalTimedEpoch(msg)
	default:
		return nil, errors.New("unhandled topic from P2P message")
	}
}

func unmarshalDecryptionKey(msg *p2p.Message) (message, error) {
	decryptionKeyMsg := shmsg.DecryptionKey{}
	if err := proto.Unmarshal(msg.Message, &decryptionKeyMsg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal decryption key P2P message")
	}

	key := new(shcrypto.EpochSecretKey)
	if err := key.Unmarshal(decryptionKeyMsg.Key); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal decryption key P2P message")
	}

	encodedKey, err := key.GobEncode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode decryption key")
	}
	return &decryptionKey{
		InstanceID: decryptionKeyMsg.InstanceID,
		EpochID:    decryptionKeyMsg.EpochID,
		Key:        encodedKey,
	}, nil
}

func unmarshalEonPublicKey(msg *p2p.Message) (message, error) {
	eonKeyMsg := shmsg.EonPublicKey{}
	if err := proto.Unmarshal(msg.Message, &eonKeyMsg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal eon public key P2P message")
	}
	return (*eonPublicKey)(&eonKeyMsg), nil
}

func unmarshalTimedEpoch(msg *p2p.Message) (message, error) {
	timedEpochMsg := shmsg.TimedEpoch{}
	if err := proto.Unmarshal(msg.Message, &timedEpochMsg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal timed epoch P2P message")
	}
	return (*timedEpoch)(&timedEpochMsg), nil
}

type unhandledTopicError struct {
	topic string
	msg   string
}

func (e *unhandledTopicError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.topic)
}
