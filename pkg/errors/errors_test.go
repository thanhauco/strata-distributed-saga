package errors

import "testing"

func TestIsRetryable(t *testing.T) {
	transient := &SagaError{Code: "TIMEOUT", Category: CategoryTransientError}
	if !IsRetryable(transient) {
		t.Errorf("Expected transient error to be retryable")
	}

	business := &SagaError{Code: "INSUFFICIENT_FUNDS", Category: CategoryBusinessError}
	if IsRetryable(business) {
		t.Errorf("Expected business error NOT to be retryable")
	}
}
