package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

const SESSION_ID_LEN = 32

func GenerateSessionId() (string, error) {
	b := make([]byte, SESSION_ID_LEN)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
