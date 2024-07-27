package domains

import "github.com/google/uuid"

type Message struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Key  uuid.UUID `json:"key"`
}
