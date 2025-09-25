package plumbing

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashSha1(content []byte) string {
	hasher := sha1.New()
	hasher.Write(content)
	b := hasher.Sum(nil)

	return hex.EncodeToString(b)
}

func IsSha1(text string) bool {
	return len(text) == sha1.Size
}
