package repository

import (
	"mgit/internal/objects"
	"os"
)

func (repo *Repository) Checkout(hash string) (string, error) {
	headHash := repo.RevParse("HEAD")
	head := objects.ParseCommit(headHash, repo.CatFile(headHash))
	commit := objects.ParseCommit(hash, repo.CatFile(hash))

	headTree := objects.ParseTree(head.Tree, repo.CatFile(head.Tree))
	targetTree := objects.ParseTree(commit.Tree, repo.CatFile(commit.Tree))

	diffs := headTree.Diff(*targetTree)

	for _, diff := range diffs {
		switch diff.Type {
		case objects.DiffDelete:
			os.Remove(diff.Path)
		case objects.DiffModified, objects.DiffInsert, objects.DiffEqual:
			content, _ := repo.storage.Get(diff.TargetHash)
			os.WriteFile(diff.Path, content, 0644)
		}
	}

	repo.UpdateHeadPointer(hash)

	return hash, nil
}
