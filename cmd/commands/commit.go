package commands

import (
	"fmt"
	"mgit/cmd/stage"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
)

func Commit(message string) string {
	stageManager := stage.Load()

	if !stageManager.HasStages() {
		fmt.Println("Nothing to commit")
		return ""
	}

	stages := stageManager.Objects()

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

	// remove deleted blobs from tree ref
	newTree.Save()

	newCommit := commit.CreateCommit(message, parentHash, newTree)
	newCommit.Save()

	stage.Truncate()
	head.UpdateHead(newCommit.Hash)

	fmt.Printf("Committed %v files\n\n", len(stages))

	return newCommit.Hash
}
