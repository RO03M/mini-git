package repository

import (
	"log"
	"os"
	"path/filepath"
)

func (repo *Repository) BranchUpdate(branch string, hash string) {
	file, err := os.OpenFile(filepath.Join(repo.DotPath, "refs/heads", branch), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("branch %s doesn't exist", branch)
		}

		log.Fatal(err)
	}

	_, err = file.WriteString(hash)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func (repo *Repository) BranchCreate(branch string) (string, error) {
	if repo.BranchExists(branch) {
		log.Fatalf("branch %s already exists", branch)
	}

	headHash := repo.RevParse("HEAD")

	err := os.WriteFile(filepath.Join(repo.DotPath, "refs/heads", branch), []byte(headHash), 0644)

	if err != nil {
		return "", err
	}

	return headHash, nil
}

func (repo *Repository) BranchExists(branch string) bool {
	stat, _ := os.Stat(filepath.Join(repo.DotPath, "refs/heads", branch))

	return stat != nil
}

func (repo *Repository) BranchGet(branch string) string {
	file, err := os.ReadFile(filepath.Join(repo.DotPath, "refs/heads", branch))

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("branch %s doesn't exist", branch)
		}

		log.Fatal(err)
	}

	return string(file)
}
