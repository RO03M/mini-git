package commands

import (
	"fmt"
	"mgit/internal/repository"
)

func RevParse(args ...string) {
	if len(args) == 0 {
		return
	}

	rev := args[0]

	repo := repository.Open()

	hash := repo.RevParse(rev)

	fmt.Println(hash)
}
