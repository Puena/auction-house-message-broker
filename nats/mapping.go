package nats

import (
	"time"

	broker "github.com/Puena/auction-house-message-broker"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// newNatsMsgFromBrokerMsg creates a new nats.Msg from a broker.Msg.
func newNatsMsgFromBrokerMsg(subject string, brokerMsg broker.Msg) *nats.Msg {
	natsMsg := nats.NewMsg(subject)
	natsMsg.Data = brokerMsg.Data
	natsMsg.Header.Add(nats.MsgIdHdr, brokerMsg.ID)
	natsMsg.Header.Add(broker.MSG_OCCURRED_AT, brokerMsg.OccurredAtToRFC3339())

	return natsMsg
}

// newBrokerMsgFromNatsMsg creates a new broker.Msg from a nats.Msg.
func newBrokerMsgFromNatsMsg(natsMsg *nats.Msg) broker.Msg {
	brokerMsg := broker.NewMsg(natsMsg.Header.Get(nats.MsgIdHdr), natsMsg.Subject, natsMsg.Data)
	if err := brokerMsg.SetOccurredAtFromRFC3339(natsMsg.Header.Get(broker.MSG_OCCURRED_AT)); err != nil {
		brokerMsg.SetOccurredAt(time.Time{})
	}

	return brokerMsg
}

// NewBrokerMsgFromNatsMsg creates a new broker.Msg from a nats.Msg.
func newBrokerPubAckFromNatsPubAck(natsPubAck *jetstream.PubAck) broker.PubAck {
	return broker.PubAck{
		Stream:    natsPubAck.Stream,
		Sequence:  natsPubAck.Sequence,
		Duplicate: natsPubAck.Duplicate,
		Domain:    natsPubAck.Domain,
	}
}

// brokerPubAckFuture is a wrapper around nats.PubAckFuture.
type brokerPubAckFuture struct {
	natsPubAckFuture jetstream.PubAckFuture
}

// newBrokerPubAckFutureFromNatsPubAckFuture creates a new brokerPubAckFuture from a nats.PubAckFuture.
func newBrokerPubAckFutureFromNatsPubAckFuture(natsPubAckFuture jetstream.PubAckFuture) *brokerPubAckFuture {
	return &brokerPubAckFuture{
		natsPubAckFuture: natsPubAckFuture,
	}
}

// Ok implements broker.PubAckFuturer.
func (a *brokerPubAckFuture) Ok() <-chan *broker.PubAck {
	transform := make(chan *broker.PubAck)
	go func() {
		for ack := range a.natsPubAckFuture.Ok() {
			brokerAck := newBrokerPubAckFromNatsPubAck(ack)
			transform <- &brokerAck
		}
		close(transform)
	}()

	return transform
}

// Err implements broker.PubAckFuturer.
func (a *brokerPubAckFuture) Err() <-chan error {
	return a.natsPubAckFuture.Err()
}

// Msg implements broker.PubAckFuturer.
func (a *brokerPubAckFuture) Msg() broker.Msg {
	return newBrokerMsgFromNatsMsg(a.natsPubAckFuture.Msg())
}
