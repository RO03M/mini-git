package commands

import (
	"errors"
	"mgit/cmd/storage"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
	"os"
)

func Checkout(ref string) error {
	currentCommit := commit.GetCommitFromHead()

	if currentCommit == nil {
		return errors.New("invalid HEAD ref, nil found")
	}

	refHash := storage.GetHashFromRef(ref)

	target := commit.FromHash(refHash)

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

	return nil
}
