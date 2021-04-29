package overlay

import (
	"encoding/json"
)

// Event is a json encoded message that gets sent to the
// frontend.
type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func NewEvent(eventType string, payload interface{}) *Event {
	return &Event{
		Type:    eventType,
		Payload: payload,
	}
}

func (e *Event) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
