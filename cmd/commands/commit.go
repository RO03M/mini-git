package commands

import (
	"fmt"
	"mgit/cmd/storage"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
)

func Commit(message string) {
	stages := storage.GetStages()

	var parentHash string
	lastCommit := commit.GetCommitFromHead()

	if lastCommit != nil {
		parentHash = lastCommit.Hash
	}

	if len(stages) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	blobs := blob.StageObjectsToBlobs(stages)

	tree := tree.CreateTree(blobs)
	tree.Save()

	commit := commit.CreateCommit(message, parentHash, tree)
	commit.Save()

	storage.ClearStage()
	head.UpdateHead(commit.Hash)

	fmt.Printf("Committed %v files\n\n", len(stages))
	fmt.Println(commit.Stringify())
}
