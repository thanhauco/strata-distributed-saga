package idempotency

import "time"

type RecordStatus string

const (
	StatusProcessing RecordStatus = "PROCESSING"
	StatusCompleted  RecordStatus = "COMPLETED"
	StatusFailed     RecordStatus = "FAILED"
)

type Record struct {
	Key       string       `json:"key"`
	Status    RecordStatus `json:"status"`
	Response  string       `json:"response"`
	CreatedAt time.Time    `json:"created_at"`
	ExpiresAt int64        `json:"expires_at"`
}

func DefaultTTL() time.Duration {
	return 24 * time.Hour
}
