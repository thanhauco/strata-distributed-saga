package resilience

import (
	"fmt"
	"testing"
	"time"
)

func TestCircuitBreakerTrips(t *testing.T) {
	cb := NewCircuitBreaker(3, 50*time.Millisecond)
	failFn := func() error { return fmt.Errorf("network timeout") }

	cb.Execute(failFn)
	cb.Execute(failFn)
	cb.Execute(failFn)

	err := cb.Execute(failFn)
	if err.Error() != "CIRCUIT_BREAKER_OPEN" {
		t.Errorf("Expected circuit breaker to be OPEN, got %v", err)
	}
}
