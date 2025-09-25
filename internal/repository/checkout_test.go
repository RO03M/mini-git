package repository_test

import (
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestModifyFileAndGoBack(t *testing.T) {
	testutils.ChDirToTemp(t)
	repo := repository.Initialize(".")

	os.WriteFile("file", []byte("v1"), 0644)
	os.WriteFile("file2", []byte("v1"), 0644)

	repo.Add("file")
	repo.Add("file2")

	c1 := repo.Commit("first commit v1")

	os.WriteFile("file", []byte("v2"), 0644)

	repo.Add("file")

	repo.Commit("second commit")

	repo.Checkout(c1.Hash)

	file, _ := os.ReadFile("file")
	if string(file) != "v1" {
		t.Fatalf("file is not in it's old version\nexpected: v1\ngot: %s", string(file))
	}
}

func TestWithDeletedFiles(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file1", []byte{}, 0644)
	os.WriteFile("file2", []byte{}, 0644)

	repo.Add("file1", "file2")

	c1 := repo.Commit("first commit")

	os.Remove("file2")

	repo.Add("file2")

	c2 := repo.Commit("removed file2")

	if c2.Hash == "" {
		t.Fatal("no changes were detected when commiting, but a file was deleted")
	}

	repo.Checkout(c1.Hash)

	if _, err := os.Stat("file2"); os.IsNotExist(err) {
		t.Fatal("Checkout should have restored file2")
	}

	repo.Checkout(c2.Hash)
	if _, err := os.Stat("file2"); err == nil {
		t.Fatal("Checkout 2 should have deleted the \"file2\"")
	}
}
