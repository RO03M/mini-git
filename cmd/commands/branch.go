package commands

import (
	"fmt"
	"log"
	"mgit/cmd/storage"
	"mgit/cmd/structures/head"
)

func CreateBranch(name string) {
	headHash, _ := head.GetHeadHash()

	err := storage.CreateBranch(name, headHash)

	if err != nil {
		log.Fatalf("failed to create a branch: %v", err)
	}

	if headHash == "" {
		fmt.Printf("Created branch %s\n", name)
	} else {
		fmt.Printf("Created branch %s with HEAD at %s\n", name, headHash)
	}

	head.UpdateHead("refs/heads/" + name)
}
