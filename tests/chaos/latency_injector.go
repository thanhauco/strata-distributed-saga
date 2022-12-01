package chaos

import (
	"math/rand"
	"time"
)

func InjectNetworkJitter(maxDelayMs int) {
	if maxDelayMs <= 0 {
		return
	}
	d := rand.Intn(maxDelayMs)
	time.Sleep(time.Duration(d) * time.Millisecond)
}
