package resilience

import (
	"testing"
	"time"
)

func TestFullJitterBounds(t *testing.T) {
	base := 100 * time.Millisecond
	max := 2 * time.Second

	for i := 0; i < 20; i++ {
		delay := FullJitter(i, base, max)
		if delay > max {
			t.Errorf("Delay %v exceeded max %v", delay, max)
		}
	}
}
