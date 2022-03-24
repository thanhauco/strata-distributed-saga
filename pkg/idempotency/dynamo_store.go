package idempotency

import (
	"context"
	"time"
)

type DynamoStore struct {
	tableName string
}

func NewDynamoStore(tableName string) *DynamoStore {
	return &DynamoStore{tableName: tableName}
}

func (d *DynamoStore) KeyHash(sagaID, stepName string) string {
	return sagaID + ":" + stepName
}

// Interface compliance check
var _ = GenerateKey
