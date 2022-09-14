package Services

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(str string) string {
	sh := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sh[:])
}
