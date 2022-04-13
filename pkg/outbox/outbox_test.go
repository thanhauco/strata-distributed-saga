package outbox

import (
	"context"
	"testing"
	"time"
)

func TestStreamForwarder(t *testing.T) {
	pub := &MockEventBridgePublisher{}
	forwarder := NewStreamForwarder(pub)

	rec := &OutboxRecord{
		EventID:     "evt-101",
		AggregateID: "order-42",
		EventType:   "OrderCreated",
		Payload:     `{"amount": 150.00}`,
		CreatedAt:   time.Now(),
		Processed:   false,
	}

	err := forwarder.ProcessStreamRecords(context.Background(), []*OutboxRecord{rec})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !rec.Processed {
		t.Errorf("Expected record to be marked processed")
	}
	if len(pub.PublishedEvents) != 1 {
		t.Errorf("Expected 1 published event")
	}
}
