package nats

import "github.com/nats-io/nats.go/jetstream"

type NatsJetStream struct {
	js jetstream.JetStream
}
