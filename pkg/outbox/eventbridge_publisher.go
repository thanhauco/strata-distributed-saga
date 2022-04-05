package outbox

import (
	"context"
	"fmt"
)

type MockEventBridgePublisher struct {
	PublishedEvents []string
}

func (p *MockEventBridgePublisher) Publish(ctx context.Context, record *OutboxRecord) error {
	p.PublishedEvents = append(p.PublishedEvents, record.EventID)
	fmt.Printf("[EventBridge] Published event: %s type: %s\n", record.EventID, record.EventType)
	record.Processed = true
	return nil
}
