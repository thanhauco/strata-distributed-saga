package idempotency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type MemoryStore struct {
	mu      sync.RWMutex
	records map[string]*Record
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{records: make(map[string]*Record)}
}

func (m *MemoryStore) Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.records[key]; exists {
		return false, nil
	}
	m.records[key] = &Record{
		Key:       key,
		Status:    StatusProcessing,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl).Unix(),
	}
	return true, nil
}

func (m *MemoryStore) Complete(ctx context.Context, key, response string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	rec, exists := m.records[key]
	if !exists {
		return fmt.Errorf("record not found")
	}
	rec.Status = StatusCompleted
	rec.Response = response
	return nil
}
