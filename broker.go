package broker

import "context"

type brokerProvider interface {
	Publisher
	Subscriber
}

type MessageBroker struct {
	provider brokerProvider
}

func NewBroker(provider brokerProvider) *MessageBroker {
	return &MessageBroker{provider: provider}
}

// PublishMsg performs a synchronous publish to a stream and waits for ack from server
// It accepts subject name (which must be bound to a stream) and nats.Message
func (b *MessageBroker) PublishMsg(ctx context.Context, msg Msg, opts ...PublishOption) (PubAck, error) {
	return b.provider.PublishMsg(ctx, msg, opts...)
}

// PublishAsync performs an asynchronous publish to a stream and returns [PubAckFuture] interface
// It accepts subject name (which must be bound to a stream) and message data
func (b *MessageBroker) PublishMsgAsync(msg Msg, opts ...PublishOption) (PubAckFuturer, error) {
	return b.provider.PublishMsgAsync(msg, opts...)
}
