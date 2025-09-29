package index_test

import (
	"mgit/internal/index"
	"mgit/internal/testutils"
	"os"
	"strings"
	"testing"
)

func TestAddRepeated(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.WriteFile("index", []byte{}, 0644)
	storage := index.Open("index")

	storage.Add("file", "hash")
	storage.Add("file", "hash")

	storage.WriteBuffer()

	file, _ := os.ReadFile("index")

	if strings.Trim(string(file), "\n") != "add hash file" {
		t.Fatalf("wrong index file content.\ngot: %s", string(file))
	}
}

func TestAddAndReopen(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.WriteFile("index", []byte{}, 0644)
	storage := index.Open("index")

	storage.Add("file", "hash")
	storage.Add("file2", "hash")

	storage.WriteBuffer()

	storage = index.Open("index")

	if len(storage.Items) != 2 {
		t.Fatalf("storage should have 2 items, but got: %v", len(storage.Items))
	}
}
