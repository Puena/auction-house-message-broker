package nats

import "github.com/nats-io/nats.go/jetstream"

type natsJetStream struct {
	js jetstream.JetStream
}

// NewNatsJetStreamProvider create a new nats jetstream provider for broker.
func NewNatsJetStreamProvider(js jetstream.JetStream) *natsJetStream {
	return &natsJetStream{js: js}
}
