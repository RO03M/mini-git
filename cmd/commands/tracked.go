package commands

import (
	"fmt"
	"log"
	"mgit/cmd/structures/commit"
)

func Tracked() {
	lastCommit := commit.GetCommitFromHead()

	if lastCommit == nil {
		fmt.Println("No commits were ever made")
		return
	}

	tree := lastCommit.Tree

	if tree == nil {
		log.Fatal("no tree attached to the commit")
		return
	}

	tree.LoadBlobs()

	blobs := tree.Blobs

	for _, blob := range blobs {
		fmt.Println(blob.FilePath)
	}

	fmt.Printf("\nTotal of %v tracked files\n", len(blobs))
}
