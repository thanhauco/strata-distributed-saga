package idempotency

import "testing"

func TestGenerateKey(t *testing.T) {
	k1 := GenerateKey("saga-1", "reserve-payment", "order-100")
	k2 := GenerateKey("saga-1", "reserve-payment", "order-100")
	k3 := GenerateKey("saga-1", "reserve-inventory", "order-100")

	if k1 != k2 {
		t.Errorf("Deterministic keys must match")
	}
	if k1 == k3 {
		t.Errorf("Different steps must yield different keys")
	}
}
