package testing

import "encoding/json"

// QueueEntry is a test struct.
type QueueEntry struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// Decode deserializes a queue entry.
func (e *QueueEntry) Decode(data []byte) error {
	err := json.Unmarshal(data, e)
	if err != nil {
		return err
	}
	return nil
}

// Encode serializes a queue entry.
func (e *QueueEntry) Encode() (data []byte, err error) {
	return json.Marshal(e)
}
