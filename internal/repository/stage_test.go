package repository_test

import (
	"fmt"
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestSimpleStage(t *testing.T) {
	testutils.ChDirToTemp(t)
	repo := repository.Initialize(".")

	os.WriteFile("file", []byte{}, 0644)

	repo.Add("file")

	if len(repo.Status().Staged) != 1 {
		t.Fatal("wrong staged files size")
	}

	if repo.Status().Staged[0].Path != "file" {
		t.Fatalf("wrong staged file path.\nexpected: file\ngot: %s", repo.Status().Staged[0].Path)
	}
}

func TestStageFilesOutOfRoot(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.MkdirAll("1/2/3/4", 0755)
	os.WriteFile("1/2/3/4/file", []byte{}, 0644)

	repo := repository.Initialize(".")

	os.Chdir("1/2/3/4")

	repo.Add("file")

	if repo.Status().Staged[0].Path != "1/2/3/4/file" {
		t.Fatalf("wrong staged file path.\nexpected: 1/2/3/4/file\ngot: %s", repo.Status().Staged[0].Path)
	}
}

func TestShouldNotStageUnmodified(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file", []byte("content"), 0644)
	repo.Add("file")
	repo.Commit("first")

	repo.Add("file")

	staged := repo.Status().Staged
	fmt.Println(staged)
	if len(staged) != 0 {
		t.Fatal("an unmodified file was added to stage")
	}
}
