package commands

import (
	"log"
	"mgit/internal/repository"
)

func Rm(paths ...string) {
	if len(paths) == 0 {
		log.Fatal("no path was given to remove, nothing changed")
	}

	repo := repository.Open()

	repo.Remove(paths...)
}
