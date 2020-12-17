package overlay

// Event is a json encoded message that gets sent to the
// frontend.
type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
