package types

type SagaState string

const (
	SagaPending     SagaState = "PENDING"
	SagaRunning     SagaState = "RUNNING"
	SagaCompleted   SagaState = "COMPLETED"
	SagaCompensating SagaState = "COMPENSATING"
	SagaFailed      SagaState = "FAILED"
)

type SagaInstance struct {
	SagaID       string    `json:"saga_id"`
	Type         string    `json:"type"`
	State        SagaState `json:"state"`
	Payload      string    `json:"payload"`
	CreatedAt    int64     `json:"created_at"`
	UpdatedAt    int64     `json:"updated_at"`
	ExpiresAt    int64     `json:"expires_at"`
}
