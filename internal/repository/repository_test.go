package repository_test

import (
	"log"
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"path/filepath"
	"testing"
)

func TestRepoInit(t *testing.T) {
	testutils.ChDirToTemp(t)
	os.Mkdir("dir", 0755)
	repository.Initialize("dir")
	_, err := os.ReadDir("dir/.mgit")

	if err != nil {
		log.Fatal(err)
	}
}

func TestOpenRepository(t *testing.T) {
	testutils.ChDirToTemp(t)

	repository.Initialize(".")

	repo := repository.Open()

	if repo == nil {
		log.Fatal("failed to open repo, <nil> returned")
	}
}

func TestOpenRepositoryInDirBelow(t *testing.T) {
	testutils.ChDirToTemp(t)
	repository.Initialize(".")
	os.Mkdir("below", 0755)
	os.Chdir("below")

	repo := repository.Open()

	cwd, _ := os.Getwd()
	relpath, _ := filepath.Rel(cwd, repo.DotPath)

	if relpath != "../.mgit" {
		log.Fatalf("wrong rel path.\nexpected: ../.mgit\ngot: %s", relpath)
	}
}

func TestRelativePath(t *testing.T) {
	testutils.ChDirToTemp(t)

	repository.Initialize(".")

	os.MkdirAll("1/2/3", 0755)
	os.WriteFile("1/2/3/file", nil, 0644)

	repo := repository.Open()

	relpath := repo.PathFromDot("1/2/3/file")

	if relpath != "1/2/3/file" {
		log.Fatalf("wrong relpath.\ngot: %s\nwant: 1/2/3/file", relpath)
	}

	os.Chdir("1/2/3")

	repo = repository.Open()

	relpath = repo.PathFromDot("file")

	if relpath != "1/2/3/file" {
		log.Fatalf("wrong relpath.\ngot: %s\nwant: 1/2/3/file", relpath)
	}

}
