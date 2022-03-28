package outbox

import "time"

type OutboxRecord struct {
	EventID     string    `json:"event_id"`
	AggregateID string    `json:"aggregate_id"`
	EventType   string    `json:"event_type"`
	Payload     string    `json:"payload"`
	CreatedAt   time.Time `json:"created_at"`
	Processed   bool      `json:"processed"`
	Version     int       `json:"version"`
}
