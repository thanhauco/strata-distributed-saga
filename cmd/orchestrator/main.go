package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/thanhauco/strata-distributed-saga/pkg/events"
)

type StartSagaRequest struct {
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	SKU        string  `json:"sku"`
	Quantity   int     `json:"quantity"`
}

func HandleStartSaga(ctx context.Context, req StartSagaRequest) (string, error) {
	sagaID := fmt.Sprintf("saga-%s", uuid.New().String()[:8])
	cmd := events.SagaCommand{
		SagaID:         sagaID,
		OrderID:        req.OrderID,
		CustomerID:     req.CustomerID,
		Amount:         req.Amount,
		SKU:            req.SKU,
		Quantity:       req.Quantity,
		IdempotencyKey: fmt.Sprintf("%s-init", sagaID),
	}
	bytes, _ := json.Marshal(cmd)
	fmt.Printf("[Orchestrator] Starting AWS Step Function execution with input: %s\n", string(bytes))
	return sagaID, nil
}

func main() {
	lambda.Start(HandleStartSaga)
}
