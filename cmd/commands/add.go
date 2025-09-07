package commands

import (
	"fmt"
	"log"
	"mgit/cmd/stage"
	"mgit/cmd/storage"
	"mgit/cmd/structures/blob"
	"os"
)

func Add(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s doesn't exist\n", path)
		return
	}

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	blob := blob.CreateBlob(file)
	storage.Create(blob.Hash, blob.Content)

	stage.AddFile(path, blob.Hash)
}
