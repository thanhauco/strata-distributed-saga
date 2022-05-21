package main

import (
	"context"
	"fmt"
	"github.com/thanhauco/strata-distributed-saga/pkg/payment"
)

var paySvc = payment.NewService()

func CompensateRefund(ctx context.Context, sagaID string) error {
	fmt.Printf("[PaymentCompensation] Rolling back authorization for Saga: %s\n", sagaID)
	return paySvc.RefundPayment(ctx, sagaID)
}
