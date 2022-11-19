package main

import (
	"context"
	"fmt"
	"time"
)

type Reconciler struct {
	stuckThreshold time.Duration
}

func NewReconciler(thresh time.Duration) *Reconciler {
	return &Reconciler{stuckThreshold: thresh}
}

func (r *Reconciler) ScanAndRepair(ctx context.Context) error {
	fmt.Println("[Reconciler] Scanning DynamoDB for RUNNING sagas past SLA threshold...")
	return nil
}
