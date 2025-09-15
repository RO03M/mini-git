package commands_test

import (
	"fmt"
	"mgit/cmd/commands"
	"mgit/cmd/structures/commit"
	"mgit/internal/testutils"
	"os"
	"strings"
	"testing"
)

func TestSimpleCommit(t *testing.T) {
	testutils.ChDirToTemp(t)
	commands.Init()

	os.WriteFile("file", []byte("content"), 0644)
	os.WriteFile("file2", []byte("content"), 0644)

	commands.Add("file")
	commands.Add("file2")

	commands.Commit("commit message")

	lastCommit := commit.GetCommitFromHead()

	if lastCommit == nil {
		t.Fatal("commit was not made")
	}

	if strings.ReplaceAll(lastCommit.Message, "\n", "") != "commit message" {
		t.Fatalf("wrong commit message\nexpected: commit message\ngot: %s", lastCommit.Message)
	}
}

func TestCommitWithDeletedFile(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.WriteFile("file", []byte("arquivo que precisa ser removido"), 0644)

	commands.Init()
	commands.Add("file")
	h1 := commands.Commit("first")

	commit.GetCommitFromHead()

	// fmt.Println("content", lastCommit.Blobs()[0].Content)
	os.Remove("file")

	commands.Add("file")
	h2 := commands.Commit("removed file")

	c1 := commit.FromHash(h1)
	c2 := commit.FromHash(h2)

	fmt.Println(c1.Blobs())
	fmt.Println(c2.Blobs())
	// firstCommit := commit.FromHash(firstCommitHash)
	// firstCommit.Tree.LoadBlobs()

	// if lastCommit.
	// fmt.Println(firstCommit.Tree.Blobs)
	// fmt.Println(lastCommit.Tree.Blobs)
}
