package payment

import (
	"context"
	"testing"
)

func TestPaymentReserveAndRefund(t *testing.T) {
	svc := NewService()
	ctx := context.Background()

	err := svc.ReservePayment(ctx, "saga-42", 500.0)
	if err != nil {
		t.Fatalf("Reserve failed: %v", err)
	}

	err = svc.RefundPayment(ctx, "saga-42")
	if err != nil {
		t.Fatalf("Refund failed: %v", err)
	}
}
