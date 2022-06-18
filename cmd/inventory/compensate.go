package main

import (
	"context"
	"fmt"
	"github.com/thanhauco/strata-distributed-saga/pkg/inventory"
)

var stockSvc = inventory.NewStockService()

func CompensateRelease(ctx context.Context, sku string, qty int) {
	fmt.Printf("[InventoryCompensation] Releasing SKU: %s, Qty: %d back to warehouse\n", sku, qty)
	stockSvc.Release(ctx, sku, qty)
}
