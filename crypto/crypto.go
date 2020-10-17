package crypto

import (
	"crypto/sha256"
)

func MainHash(data []byte) [32]byte {
	return sha256.Sum256(data)
}
