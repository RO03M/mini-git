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
