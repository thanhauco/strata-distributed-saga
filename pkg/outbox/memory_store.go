package outbox

import (
	"context"
	"sync"
)

type MemoryOutbox struct {
	mu      sync.Mutex
	records []*OutboxRecord
}

func NewMemoryOutbox() *MemoryOutbox {
	return &MemoryOutbox{records: make([]*OutboxRecord, 0)}
}

func (m *MemoryOutbox) Save(ctx context.Context, rec *OutboxRecord) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.records = append(m.records, rec)
	return nil
}

func (m *MemoryOutbox) GetUnprocessed() []*OutboxRecord {
	m.mu.Lock()
	defer m.mu.Unlock()
	var unproc []*OutboxRecord
	for _, r := range m.records {
		if !r.Processed {
			unproc = append(unproc, r)
		}
	}
	return unproc
}
