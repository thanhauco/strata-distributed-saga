package resilience

import (
	"math"
	"math/rand"
	"time"
)

// FullJitter calculates backoff delay using the AWS Architecture Full Jitter formula:
// Sleep = rand(0, min(max_delay, base_delay * 2^attempt))
func FullJitter(attempt int, baseDelay, maxDelay time.Duration) time.Duration {
	temp := float64(baseDelay) * math.Pow(2, float64(attempt))
	capped := math.Min(float64(maxDelay), temp)
	return time.Duration(rand.Float64() * capped)
}
