package outbox

import (
	"context"
)

type StreamForwarder struct {
	publisher *MockEventBridgePublisher
}

func NewStreamForwarder(pub *MockEventBridgePublisher) *StreamForwarder {
	return &StreamForwarder{publisher: pub}
}

func (f *StreamForwarder) ProcessStreamRecords(ctx context.Context, records []*OutboxRecord) error {
	for _, rec := range records {
		if err := f.publisher.Publish(ctx, rec); err != nil {
			return err
		}
	}
	return nil
}
