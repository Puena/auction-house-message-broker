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

func (b *MessageBroker) PublishMsg(ctx context.Context, subject string, msg Msg, opts ...PublishOption) (PubAck, error) {
	return b.provider.PublishMsg(ctx, subject, msg, opts...)
}

func (b *MessageBroker) PublishMsgAsync(subject string, msg Msg, opts ...PublishOption) (PubAckFuturer, error) {
	return b.provider.PublishMsgAsync(subject, msg, opts...)
}
