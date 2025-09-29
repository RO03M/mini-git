package commands

import (
	"fmt"
	"log"
	"mgit/internal/repository"
)

func Switch(args ...string) {
	if len(args) == 0 {
		log.Fatal("missing branch name")
	}

	if len(args) > 1 {
		log.Fatal("too many arguments, expected only one")
	}

	repo := repository.Open()

	repo.Switch(args[0])

	fmt.Printf("Switched to branch \"%s\"\n", args[0])
}
