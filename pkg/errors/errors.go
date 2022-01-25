package errors

import "fmt"

type ErrorCategory string

const (
	CategoryTransientError ErrorCategory = "TRANSIENT"
	CategoryBusinessError  ErrorCategory = "BUSINESS"
	CategoryFatalError     ErrorCategory = "FATAL"
)

type SagaError struct {
	Code     string        `json:"code"`
	Message  string        `json:"message"`
	Category ErrorCategory `json:"category"`
}

func (e *SagaError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Category, e.Code, e.Message)
}

func IsRetryable(err error) bool {
	if se, ok := err.(*SagaError); ok {
		return se.Category == CategoryTransientError
	}
	return false
}
