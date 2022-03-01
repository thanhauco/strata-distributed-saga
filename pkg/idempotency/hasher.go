package idempotency

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateKey(components ...string) string {
	h := sha256.New()
	for _, c := range components {
		h.Write([]byte(c))
		h.Write([]byte("|"))
	}
	return hex.EncodeToString(h.Sum(nil))
}
