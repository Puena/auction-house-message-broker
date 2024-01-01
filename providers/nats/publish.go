package nats

import (
	"context"

	broker "github.com/Puena/auction-house-message-broker"
	"github.com/nats-io/nats.go/jetstream"
)

// Publish publishes a message to the nats jetstream broker synchronous.
func (n *natsJetStream) Publish(ctx context.Context, msg broker.Msg, opts ...broker.PublishOption) (broker.PubAck, error) {
	options, err := broker.NewPublishOptions(opts...)
	if err != nil {
		return broker.PubAck{}, err
	}

	natsOpts := make([]jetstream.PublishOpt, 0)

	if options.MsgRetry != 0 {
		natsOpts = append(natsOpts, jetstream.WithRetryAttempts(options.MsgRetry))
	}

	natsMsg := newNatsMsgFromBrokerMsg(msg)
	ack, err := n.js.PublishMsg(ctx, natsMsg, natsOpts...)
	if err != nil {
		return broker.PubAck{}, err
	}

	return newBrokerPubAckFromNatsPubAck(ack), nil
}

// PublishAsync publishes a message to the nats jetstream broker asynchronous.
func (n *natsJetStream) PublishAsync(msg broker.Msg, opts ...broker.PublishOption) (broker.PubAckFuturer, error) {
	options, err := broker.NewPublishOptions(opts...)
	if err != nil {
		return nil, err
	}

	natsOpts := make([]jetstream.PublishOpt, 0, len(opts))

	if options.MsgRetry != 0 {
		natsOpts = append(natsOpts, jetstream.WithRetryAttempts(options.MsgRetry))
	}

	natsMsg := newNatsMsgFromBrokerMsg(msg)
	ack, err := n.js.PublishMsgAsync(natsMsg, natsOpts...)
	if err != nil {
		return nil, err
	}

	return newBrokerPubAckFutureFromNatsPubAckFuture(ack), nil
}
