package commands_test

import (
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

	os.Remove("file")

	commands.Add("file")
	h2 := commands.Commit("removed file")

	c1 := commit.FromHash(h1)
	c2 := commit.FromHash(h2)

	if len(c1.Blobs()) != 1 || len(c2.Blobs()) != 0 {
		t.Fatalf("Wrong commit blobs sizes\nc1: %v\n c2: %v\n", c1.Blobs(), c2.Blobs())
	}
}
