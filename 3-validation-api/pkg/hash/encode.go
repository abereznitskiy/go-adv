package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func EncodeEmail(email string) string {
	normalizedEmail := strings.ToLower(email)
	hash := sha256.New()
	hash.Write([]byte(normalizedEmail))

	return hex.EncodeToString(hash.Sum(nil))
}
