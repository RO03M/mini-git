package storage

import (
	"log"
	"mgit/internal/plumbing"
	"os"
	"path/filepath"
)

type Hash string

type Storage struct {
	ObjectsPath string
}

func (s Storage) Create(content []byte) (Hash, error) {
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

	return Hash(hash), nil
}
