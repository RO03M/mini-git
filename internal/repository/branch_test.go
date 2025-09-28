package repository_test

import (
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestCreateBranch(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file", []byte{}, 0644)
	repo.Add("file")
	c1 := repo.Commit("initial commit")

	branch1Hash, err := repo.BranchCreate("branch1")

	if err != nil {
		t.Fatalf("couldn't create branch1: %v", err)
	}

	if branch1Hash != c1.Hash {
		t.Fatalf("wrong hash value.\nwant: %s\ngot: %v", c1.Hash, branch1Hash)
	}

	if !repo.BranchExists("branch1") {
		t.Fatal("branch1 wasn't created")
	}
}

func TestSwitchBranch(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file", []byte{}, 0644)
	repo.Add("file")
	repo.Commit("initial commit")

	repo.BranchCreate("branch1")
	repo.BranchCreate("branch2")

	repo.Switch("branch2")

	os.WriteFile("file2", []byte("file2"), 0644)
	repo.Add("file2")
	c2 := repo.Commit("second commit")

	repo.Switch("branch1")

	if stat, _ := os.Stat("file2"); stat != nil {
		t.Fatal("file2 was not deleted")
	}

	if repo.RevParse("branch2") != c2.Hash {
		t.Fatal("the branch2 hash is different from the second commit")
	}

	if repo.RevParse("branch1") != repo.RevParse("master") {
		t.Fatal("the branch1 hash is different from the master branch")
	}
}
