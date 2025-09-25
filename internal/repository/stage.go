package repository

import (
	"log"
	"os"
)

func (repo *Repository) Add(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// addRemovedFile(path)
			// return
			log.Fatal("path doesn't exist: ", path)
		}

		file, err := os.ReadFile(path)

		if err != nil {
			log.Fatal(err)
		}

		hash, err := repo.storage.Create(file)

		if err != nil {
			log.Fatal(err)
		}

		repo.index.Add(path, hash)
	}

	err := repo.index.WriteBuffer()

	if err != nil {
		log.Fatal(err)
	}
}
