package tree_test

import (
	"fmt"
	"log"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/tree"
	"testing"
)

func TestTreeDiff(t *testing.T) {
	tree1 := tree.Tree{
		Blobs: []blob.Blob{
			{
				Hash:     "a",
				FilePath: "file1",
			},
			{
				Hash:     "b",
				FilePath: "file2",
			},
			{
				Hash:     "c",
				FilePath: "file3",
			},
		},
	}

	// "file2" was deleted
	tree2 := tree.Tree{
		Blobs: []blob.Blob{
			{
				Hash:     "a", // No changes
				FilePath: "file1",
			},
			{
				Hash:     "c2", // Changed
				FilePath: "file3",
			},
			{
				Hash:     "d",
				FilePath: "file4", // New File
			},
		},
	}

	diffs := tree1.Diff(tree2)

	correctCaseMap := map[string]tree.TreeDiff{
		"file1": {
			Type: tree.DiffEqual,
		},
		"file2": {
			Type: tree.DiffDelete,
		},
		"file3": {
			Type: tree.DiffModified,
		},
		"file4": {
			Type: tree.DiffInsert,
		},
	}

	for _, diff := range diffs {
		correctCase := correctCaseMap[diff.Blob.FilePath]

		if diff.Type != correctCase.Type {
			t.Fatalf("Wrong diff type for blob %s\nexpected: %v\ngot: %v", diff.Blob.FilePath, correctCase.Type, diff.Type)
		}
	}
}

func TestMergeTreeShouldReplaceExistingBlobs(t *testing.T) {
	var blobs []blob.Blob = []blob.Blob{
		{
			FilePath: "a",
			Hash:     "old",
		},
		{
			FilePath: "b",
		},
		{
			FilePath: "c",
		},
	}

	parentTree := tree.CreateTree(blobs)

	tree := tree.CreateMergedTree(parentTree, []blob.Blob{{FilePath: "a", Hash: "new"}})

	if len(tree.Blobs) != 3 {
		log.Fatal("wrong blobs size")
	}

	fmt.Println(tree)
}
