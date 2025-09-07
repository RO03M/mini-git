package storage_test

import (
	"log"
	"mgit/cmd/storage"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	testutils.ChDirToTemp(t)
	storage.Init()
	_, err := os.ReadDir("./.mgit")

	if err != nil {
		log.Fatal(err)
	}
}
