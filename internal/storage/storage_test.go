package storage_test

import (
	"mgit/internal/storage"
	"mgit/internal/testutils"
	"os"
	"path/filepath"
	"testing"
)

func TestStorageCreate(t *testing.T) {
	testutils.ChDirToTemp(t)
	os.Mkdir("objects", 0755)

	storagePath, _ := os.Getwd()

	storagePath = filepath.Join(storagePath, "objects")

	store := storage.Storage{
		ObjectsPath: storagePath,
	}

	hash, err := store.Create([]byte("object content"))

	if err != nil {
		t.Fatalf("there was an error while creating the object in the storage: %v", err)
	}

	if hash == "" {
		t.Fatal("hash is empty")
	}

	if !store.Exists(hash) {
		t.Fatal("object was not created in the filesystem")
	}
}

func TestStorageGet(t *testing.T) {
	testutils.ChDirToTemp(t)

	abs, _ := filepath.Abs(".")

	store := storage.Storage{
		ObjectsPath: abs,
	}

	hash, _ := store.Create([]byte("content"))

	object, err := store.Get(hash)

	if err != nil {
		t.Fatal(err)
	}

	if string(object) != "content" {
		t.Fatalf("wrong object's content.\ngot: %v\nwant: %v", string(object), "content")
	}
}
