package order

import (
	"context"
	"testing"
)

func TestOrderCreation(t *testing.T) {
	svc := NewService()
	o := &Order{ID: "ord-1", CustomerID: "cust-9", Amount: 299.99}

	err := svc.CreateOrder(context.Background(), o)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if o.Status != StatusSubmitted {
		t.Errorf("Expected status SUBMITTED, got %s", o.Status)
	}
}
