package shipping

import (
	"context"
	"testing"
)

func TestShippingDispatchAndCancel(t *testing.T) {
	svc := NewService()
	ctx := context.Background()

	shipment, err := svc.Dispatch(ctx, "ord-99")
	if err != nil || shipment == nil {
		t.Fatalf("Dispatch failed")
	}

	svc.Cancel(ctx, "ord-99")
	if len(svc.shipments) != 0 {
		t.Errorf("Shipment not cancelled")
	}
}
