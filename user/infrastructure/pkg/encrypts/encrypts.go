package encrypts

import (
	"crypto/sha256"
	"encoding/hex"
)

const secret = "newsCenter-secret"

func EncryptPassword(Password string) string {
	h := sha256.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(Password)))
}
