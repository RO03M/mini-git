package repository

import (
	"fmt"
	"log"
	"mgit/internal/objects"
	"os"
)

// func (repo *Repository) ApplyDiffs(headHash string, referenceHash string) {
// 	head := repo.GetCommit(headHash)

// 	if head == nil {
// 		return "", fmt.Errorf("failed to get head commit with hash %s, nil returned", headHash)
// 	}

// 	reference := repo.GetCommit(referenceHash)

// 	if reference == nil {
// 		return "", fmt.Errorf("failed to get reference with hash %s, nil returned", referenceHash)
// 	}

// 	headTree := repo.GetTree(head.Tree)
// 	referenceTree := repo.GetTree(reference.Tree)

// 	diffs := headTree.Diff(*referenceTree)

// 	for _, diff := range diffs {
// 		switch diff.Type {
// 		case objects.DiffDelete:
// 			os.Remove(diff.Path)
// 		case objects.DiffModified, objects.DiffInsert, objects.DiffEqual:
// 			content, _ := repo.storage.Get(diff.TargetHash)
// 			os.WriteFile(diff.Path, content, 0644)
// 		}
// 	}
// }

func (repo *Repository) Checkout(hash string) (string, error) {
	head := repo.GetCommit(repo.RevParse("HEAD"))

	if head == nil {
		return "", fmt.Errorf("failed to get HEAD, nil returned")
	}

	commit := repo.GetCommit(hash)

	if commit == nil {
		return "", fmt.Errorf("failed to get commit with hash %s, nil returned", hash)
	}

	headTree := repo.GetTree(head.Tree)
	targetTree := repo.GetTree(commit.Tree)

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

	return hash, nil
}

func (repo *Repository) Switch(branch string) {
	// switch de branch tem que atualizar o head para usar o ref da branch

	if !repo.BranchExists(branch) {
		log.Fatalf("branch %s doesn't exist", branch)
	}

	if len(repo.index.Items) != 0 {
		log.Fatal("there are staged files to be committed")
	}

	repo.Checkout(repo.RevParse(branch))
	repo.UpdateHeadDirect("refs/heads/" + branch)
}
