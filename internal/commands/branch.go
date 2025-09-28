package commands

import (
	"fmt"
	"log"
	"mgit/internal/plumbing"
	"mgit/internal/repository"
)

func listBranches() {
	repo := repository.Open()

	var activeBranch string

	ref, isBranch := getCurrentBranch(*repo)

	if isBranch {
		activeBranch = ref
	}

	branches := repo.Branches()

	for _, branch := range branches {
		if branch == activeBranch {
			plumbing.PrintfColor(plumbing.ColorGreen, "%s", branch)
			fmt.Println(" *")
			continue
		}

		fmt.Println(branch)
	}
}

func createBranch(name string) {
	repo := repository.Open()

	hash, err := repo.BranchCreate(name)

	if err != nil {
		log.Fatalf("failed to create a branch: %v", err)
	}

	fmt.Printf("Created branch %s with commit %s\n", name, hash)
}

func Branch(args ...string) {
	if len(args) == 0 {
		listBranches()
		return
	}

	if len(args) == 1 {
		createBranch(args[0])
		return
	}
}
