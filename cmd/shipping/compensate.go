package main

import (
	"context"
	"fmt"
	"github.com/thanhauco/strata-distributed-saga/pkg/shipping"
)

var shipSvc = shipping.NewService()

func CompensateShipment(ctx context.Context, orderID string) {
	fmt.Printf("[ShippingCompensation] Revoking carrier manifest for order %s\n", orderID)
	shipSvc.Cancel(ctx, orderID)
}
