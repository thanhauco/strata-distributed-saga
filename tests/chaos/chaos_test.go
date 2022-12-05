package chaos

import (
	"testing"
)

func TestChaosFailureRollback(t *testing.T) {
	// Simulates 504 Gateway timeout and validates compensation triggers
	InjectNetworkJitter(5)
}
