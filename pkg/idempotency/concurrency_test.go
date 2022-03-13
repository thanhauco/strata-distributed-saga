package idempotency

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestConcurrentAcquire(t *testing.T) {
	store := NewMemoryStore()
	key := "shared-idemp-key"
	ctx := context.Background()

	var wg sync.WaitGroup
	acquiredCount := 0
	var mu sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok, _ := store.Acquire(ctx, key, time.Minute)
			if ok {
				mu.Lock()
				acquiredCount++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if acquiredCount != 1 {
		t.Fatalf("Expected exactly 1 acquisition, got %d", acquiredCount)
	}
}
