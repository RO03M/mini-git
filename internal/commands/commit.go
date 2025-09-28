package commands

import (
	"fmt"
	"log"
	"mgit/internal/repository"
)

func Commit(args ...string) {
	if len(args) == 0 {
		log.Fatal("write a message to the commit: (mgit commit \"<message>\")")
		return
	}

	var message string

	if len(args) == 1 {
		message = args[0]
	}

	repo := repository.Open()

	commit := repo.Commit(message)

	if commit.IsEmpty() {
		log.Fatal("invalid commit created")
	}

	fmt.Printf("commit %s created\n", commit.Hash)
}
