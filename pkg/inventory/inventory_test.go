package inventory

import (
	"context"
	"testing"
)

func TestStockReservation(t *testing.T) {
	svc := NewStockService()
	ctx := context.Background()

	err := svc.Reserve(ctx, "LAPTOP-PRO-16", 2)
	if err != nil {
		t.Fatalf("Unexpected reservation error: %v", err)
	}

	err = svc.Reserve(ctx, "PHONE-MAX", 1)
	if err == nil {
		t.Fatalf("Expected out of stock error for PHONE-MAX")
	}

	svc.Release(ctx, "LAPTOP-PRO-16", 2)
	if svc.stock["LAPTOP-PRO-16"] != 50 {
		t.Errorf("Stock was not correctly released back")
	}
}
