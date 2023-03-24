package p2p

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
)

// gossipRoom represents a subscription to a single PubSub topic. Messages
// can be published to the topic with gossipRoom.Publish, and received
// messages are pushed to the Messages channel.
type gossipRoom struct {
	pubSub       *pubsub.PubSub
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	topicName    string
	self         peer.ID
}

// Publish sends a message to the pubsub topic.
func (room *gossipRoom) Publish(ctx context.Context, message []byte) error {
	return room.topic.Publish(ctx, message)
}

func (room *gossipRoom) ListPeers() []peer.ID {
	return room.pubSub.ListPeers(room.topicName)
}

// readLoop pulls messages from the pubsub topic and pushes them onto the given messages channel.
func (room *gossipRoom) readLoop(ctx context.Context, messages chan *Message) error {
	for {
		msg, err := room.subscription.Next(ctx)
		if err != nil {
			return err
		}
		// only forward messages delivered by others
		if msg.ReceivedFrom == room.self {
			continue
		}
		m := &Message{
			Topic:        room.topicName,
			Message:      msg.Data,
			Sender:       msg.GetFrom(),
			ReceivedFrom: msg.ReceivedFrom,
			ID:           msg.ID,
		}
		// send valid messages onto the Messages channel
		select {
		case messages <- m:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
