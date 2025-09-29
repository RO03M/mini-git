package repository

import (
	"log"
	"mgit/internal/plumbing"
	"os"
	"path/filepath"
)

func (repo *Repository) GetHead() string {
	head, _ := os.ReadFile(filepath.Join(repo.DotPath, "HEAD"))

	return string(head)
}

// updates the head or the target that it is pointing
func (repo *Repository) UpdateHeadPointer(hash string) {
	head := repo.GetHead()

	if branch := plumbing.BranchFromRef(head); branch != "" {
		repo.BranchUpdate(branch, hash)
		return
	}

	repo.UpdateHeadDirect(hash)
}

// replaces the HEAD file contents (even if it is a pointer)
func (repo *Repository) UpdateHeadDirect(hash string) {
	file, err := os.OpenFile(filepath.Join(repo.DotPath, "HEAD"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(hash)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func (repo *Repository) RevParse(ref string) string {
	if ref == "HEAD" {
		head := repo.GetHead()

		if branch := plumbing.BranchFromRef(head); branch != "" {
			return repo.RevParse(head)
		}

		return head
	}

	if branch := plumbing.BranchFromRef(ref); branch != "" {
		return repo.BranchGet(branch)
	}

	if repo.BranchExists(ref) {
		return repo.BranchGet(ref)
	}

	return ref
}
