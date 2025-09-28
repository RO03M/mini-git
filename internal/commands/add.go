package commands

import (
	"fmt"
	"mgit/internal/repository"
)

func Add(files ...string) {
	if len(files) == 0 {
		fmt.Println("No file passed, nothing was changed")
		return
	}

	repo := repository.Open()

	repo.Add(files...)
}
