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

	if len(stages) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	var parentHash string
	var newTree *tree.Tree

	lastCommit := commit.GetCommitFromHead()

	blobs := blob.StageObjectsToBlobs(stages)

	if lastCommit == nil {
		newTree = tree.CreateTree(blobs)
	} else {
		newTree = tree.CreateMergedTree(lastCommit.Tree, blobs)
		parentHash = lastCommit.Hash
	}

	newTree.Save()

	newCommit := commit.CreateCommit(message, parentHash, newTree)
	newCommit.Save()

	storage.ClearStage()
	head.UpdateHead(newCommit.Hash)

	fmt.Printf("Committed %v files\n\n", len(stages))
}
