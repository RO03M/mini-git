package commands

import (
	"fmt"
	"mgit/internal/plumbing"
	"mgit/internal/repository"
)

// returns branch name or hash and if is a branch or a detached commit
func getCurrentBranch(repo repository.Repository) (string, bool) {
	head := repo.GetHead()

	if plumbing.IsRefPath(head) {
		return plumbing.BranchFromRef(head), true
	}

	return head, false
}

func Status() {
	repo := repository.Open()

	ref, isBranch := getCurrentBranch(*repo)

	if isBranch {
		fmt.Printf("On branch %s\n", ref)
	} else {
		fmt.Printf("Detached from any branch. On commit %s\n", ref)
	}

	status := repo.Status()

	fmt.Print("Staged files:\n(use \"mgit rm <file...>\" to unstage)\n")

	for _, staged := range status.Staged {
		plumbing.PrintfColor(plumbing.ColorGreen, "\t%s:\t%s\n", staged.Action, staged.Path)
	}

	fmt.Println()

	fmt.Print("Untracked files:\n(use \"mgit add <file...>\" to stage)\n")

	for _, path := range status.Untracked {
		plumbing.PrintfColor(plumbing.ColorRed, "\t%s\n", path)
	}
}
