package repository_test

import (
	"fmt"
	"mgit/internal/index"
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

func TestAddRemovedFile(t *testing.T) {
	testutils.ChDirToTemp(t)
	repo := repository.Initialize(".")

	os.WriteFile("file_to_be_staged", []byte("content"), 0644)
	repo.Add("file_to_be_staged")
	repo.Commit("first")

	os.Remove("file_to_be_staged")
	repo.Add("file_to_be_staged")

	staged := repo.Status().Staged

	if len(staged) == 0 {
		t.Fatal("no stages were made")
	}

	if staged[0].Path != "file_to_be_staged" {
		t.Fatalf("wrong staged file: %s", staged[0].Path)
	}

	if staged[0].Action != index.ActionDelete {
		t.Fatal("the stage action isn't a deletion")
	}
}
