package storage

import (
	"fmt"
	"log"
	"mgit/internal/plumbing"
	"os"
	"path/filepath"
)

type Storage struct {
	ObjectsPath string
}

func (s Storage) Create(content []byte) (string, error) {
	compressed := plumbing.Compress(content)
	hash := plumbing.HashSha1(content)

	objectDir := filepath.Join(s.ObjectsPath, hash[:2])
	err := os.MkdirAll(objectDir, 0755)

	if err != nil {
		log.Fatalf("failed to create the object's dir: %v", err)
	}

	err = os.WriteFile(filepath.Join(objectDir, hash[2:]), compressed, 0644)

	if err != nil {
		log.Fatalf("failed to create object: %v", err)
	}

	return hash, nil
}

func (s Storage) Get(hash string) ([]byte, error) {
	if !plumbing.IsSha1(hash) {
		return []byte{}, fmt.Errorf("\"%s\" is not a valid sha1 hash", hash)
	}

	snapshot, err := os.ReadFile(filepath.Join(s.ObjectsPath, hash[:2], hash[2:]))

	if err != nil {
		return []byte{}, nil
	}

	decompressed := plumbing.Decompress(snapshot)

	return decompressed, nil
}

func (s Storage) Exists(hash string) bool {
	if len(hash) < 4 {
		return false
	}

	stat, _ := os.Stat(filepath.Join(s.ObjectsPath, hash[:2], hash[2:]))

	return stat != nil
}
