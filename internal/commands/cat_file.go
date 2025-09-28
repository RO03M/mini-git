package commands

import (
	"fmt"
	"log"
	"mgit/internal/repository"
)

func Catfile(hash string) {
	repo := repository.Open()

	content, err := repo.CatFile(hash)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
