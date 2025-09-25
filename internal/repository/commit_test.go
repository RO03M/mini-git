package repository_test

import (
	"fmt"
	"log"
	"mgit/internal/objects"
	"mgit/internal/repository"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestSimpleCommit(t *testing.T) {
	testutils.ChDirToTemp(t)
	repo := repository.Initialize(".")

	os.WriteFile("file", []byte("content"), 0644)
	os.WriteFile("file2", []byte("content2"), 0644)

	repo.Add("file", "file2")

	commit := repo.Commit("commit message")

	if commit == nil {
		t.Fatal("commit was not created (or not returned properly, since it is <nil>)")
	}

	if commit.Hash != repo.RevParse("HEAD") {
		t.Fatal("the head wasn't updated with the new commit hash")
	}

	if len(repo.Status().Staged) != 0 {
		t.Fatal("the staged files were not cleared")
	}
}

func TestTreeIsInheriting(t *testing.T) {
	testutils.ChDirToTemp(t)

	repo := repository.Initialize(".")

	os.WriteFile("file", []byte("content"), 0644)
	repo.Add("file")
	c1 := repo.Commit("first commit")

	os.WriteFile("file2", []byte("content2"), 0644)
	repo.Add("file2")
	c2 := repo.Commit("second commit")

	tree1 := objects.ParseTree(c1.Tree, repo.CatFile(c1.Tree))
	tree2 := objects.ParseTree(c2.Tree, repo.CatFile(c2.Tree))

	if len(tree2.Entries) != 2 {
		log.Fatalf("tree2 should have 2 entries but got: %v", len(tree2.Entries))
	}
	fmt.Println(tree1.Entries, tree2.Entries)
}

func TestCommitWithDeletedFile(t *testing.T) {
	testutils.ChDirToTemp(t)
	repo := repository.Initialize(".")

	os.WriteFile("file", []byte("arquivo que precisa ser removido"), 0644)

	repo.Add("file")
	c1 := repo.Commit("first")

	os.Remove("file")

	repo.Add("file")
	c2 := repo.Commit("removed file")

	tree1 := objects.ParseTree(c1.Tree, repo.CatFile(c1.Tree))
	tree2 := objects.ParseTree(c2.Tree, repo.CatFile(c2.Tree))

	if len(tree1.Entries) != 1 || len(tree2.Entries) != 0 {
		t.Fatal("wrong tree entries size")
	}
}
