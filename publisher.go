package broker

import (
	"context"
	"errors"
)

const (
	defaultMsgRetry = 3
	MSG_OCCURRED_AT = "msg-occurred-at"
)

var (
	ErrInvalidMsgRetry = errors.New("invalid message retry number")
)

// PublishOption represent publish option.
type PublishOption func(*publishOptions)

// WithMsgRetry change default msg retry attempts, shouldn't be negative.
func WithMsgRetry(msgRetry int) PublishOption {
	return func(o *publishOptions) {
		o.MsgRetry = msgRetry
	}
}

type publishOptions struct {
	MsgRetry int
}

func NewPublishOptions(opts ...PublishOption) (publishOptions, error) {
	options := &publishOptions{
		MsgRetry: defaultMsgRetry,
	}

	options.apply(opts...)

	if err := options.validate(); err != nil {
		return publishOptions{}, err
	}

	return *options, nil
}

/* func (o *publishOptions) defaults() {
	if o.MsgRetry == 0 {
		o.MsgRetry = defaultMsgRetry
	}
} */

func (o *publishOptions) apply(opts ...PublishOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *publishOptions) validate() error {
	if o.MsgRetry < 0 {
		return ErrInvalidMsgRetry
	}

	return nil
}

type Publisher interface {
	PublishMsg(ctx context.Context, msg Msg, opts ...PublishOption) (PubAck, error)
	PublishMsgAsync(msg Msg, opts ...PublishOption) (PubAckFuturer, error)
}
