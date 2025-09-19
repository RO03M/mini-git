package commands

import (
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
	"os"
)

func Checkout(ref string) {
	currentCommit := commit.GetCommitFromHead()

	target := commit.FromHash(ref)

	target.Tree.LoadBlobs()

	diffs := currentCommit.Tree.Diff(*target.Tree)

	for _, diff := range diffs {
		switch diff.Type {
		case tree.DiffDelete:
			os.Remove(diff.TargetBlob.FilePath)
		case tree.DiffModified, tree.DiffInsert, tree.DiffEqual:
			content := diff.TargetBlob.ReadContent()
			os.WriteFile(diff.TargetBlob.FilePath, content, 0644)
		}
	}

	head.UpdateHead(target.Hash)
}
