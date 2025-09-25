package repository

import (
	"log"
	"mgit/internal/plumbing"
	"os"
)

func (repo *Repository) AddRm(paths ...string) {
	head := repo.RevParse("HEAD")
	tracked := repo.TrackedFiles(head)
	trackedMap := plumbing.StringSliceMap(tracked)

	for _, path := range paths {
		if _, found := trackedMap[path]; !found {
			log.Fatal(path + " doesn't exist")
		}

		repo.index.AddRm(path)
	}
}

func (repo *Repository) Add(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			repo.AddRm(path)
			continue
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
