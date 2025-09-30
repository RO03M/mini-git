package commands

import (
	"fmt"
	"log"
	"mgit/internal/repository"
)

func Checkout(args ...string) {
	if len(args) == 0 {
		log.Fatal("no args")
	}

	repo := repository.Open()

	ref := args[0]

	target := repo.RevParse(ref)

	if repo.BranchExists(ref) {
		repo.Switch(ref)
	} else {
		hash, err := repo.Checkout(target)

		if err != nil {
			log.Fatal(err)
		}

		repo.UpdateHeadDirect(hash)

		fmt.Printf("Switched to %s\n", hash)
	}
}
