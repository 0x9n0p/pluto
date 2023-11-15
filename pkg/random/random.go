package random

import (
	"crypto/rand"
	"encoding/hex"
)

func String(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
