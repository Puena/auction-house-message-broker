package broker

type PubAck struct {
	Stream    string `json:"stream"`
	Sequence  uint64 `json:"seq"`
	Duplicate bool   `json:"duplicate,omitempty"`
	Domain    string `json:"domain,omitempty"`
}

type PubAckFuture struct {
	ok  <-chan *PubAck
	err <-chan error
	msg *Msg
}

func NewPubAckFuture(msg *Msg) *PubAckFuture {
	return &PubAckFuture{
		ok:  make(chan *PubAck),
		err: make(chan error),
		msg: msg,
	}
}

func (a *PubAckFuture) Ok() <-chan *PubAck {
	return a.ok
}

func (a *PubAckFuture) Err() <-chan error {
	return a.err
}

func (a *PubAckFuture) Msg() *Msg {
	return a.msg
}

type PubAckFuturer interface {
	// Ok returns a receive only channel that can be used to get a PubAck.
	Ok() <-chan *PubAck

	// Err returns a receive only channel that can be used to get the error from an async publish.
	Err() <-chan error

	// Msg returns the message that was sent to the server.
	Msg() Msg
}
