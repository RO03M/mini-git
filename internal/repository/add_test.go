package repository_test

import (
	"fmt"
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestAddToIndex(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file", []byte(""), 0644)

	repo.Add("file")

	repo2 := repository.Open()

	fmt.Println(len(repo2.Status().Staged))
}
