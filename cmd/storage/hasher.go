package storage

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
)

func GenerateSha1(raw []byte) string {
	hasher := sha1.New()
	_, err := hasher.Write(raw)

	if err != nil {
		log.Fatal(err)
	}

	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}
