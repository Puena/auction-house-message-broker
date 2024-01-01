package broker

import (
	"time"
)

// Msg represents a broker message.
type Msg struct {
	ID      string
	Subject string
	Data    []byte
	// internal
	occurredAt time.Time
	// msg size from nats or another broker
	msz int
}

func NewMsg(id, subject string, data []byte) Msg {
	return Msg{
		ID:      id,
		Subject: subject,
		Data:    data,
		msz:     0,
	}
}

func (m *Msg) SetOccurredAt(t time.Time) {
	m.occurredAt = t
}

func (m *Msg) SetOccurredAtFromRFC3339(s string) error {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	m.occurredAt = t

	return nil
}

func (m *Msg) OccurredAtToRFC3339() string {
	return m.occurredAt.Format(time.RFC3339)
}

func (m *Msg) Size() int {
	return m.msz
}

func (m *Msg) OccurredAt() time.Time {
	return m.occurredAt
}
