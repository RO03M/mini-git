package commands_test

import (
	"mgit/cmd/commands"
	"mgit/cmd/structures/commit"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestModifyFileAndGoBack(t *testing.T) {
	testutils.ChDirToTemp(t)
	commands.Init()

	os.WriteFile("file", []byte("v1"), 0644)
	os.WriteFile("file2", []byte("v1"), 0644)

	commands.Add("file")
	commands.Add("file2")

	commands.Commit("first commit")

	v1Commit := commit.GetCommitFromHead()

	os.WriteFile("file", []byte("v2"), 0644)

	commands.Add("file")

	commands.Commit("second commit")

	commands.Checkout(v1Commit.Hash)

	file, _ := os.ReadFile("file")
	if string(file) != "v1" {
		t.Fatalf("file is not in it's old version\nexpected: v1\ngot: %s", string(file))
	}
}

func TestTreeShouldInherit(t *testing.T) {
	testutils.ChDirToTemp(t)
	commands.Init()

	os.WriteFile("file_first_commit", []byte("v1"), 0644)

	commands.Add("file_first_commit")

	commands.Commit("first commit")

	//

	os.WriteFile("file_second_commit", []byte("v2"), 0644)

	commands.Add("file_second_commit")

	commands.Commit("second commit")

	secondCommit := commit.GetCommitFromHead()

	os.Remove("file_first_commit")
	os.Remove("file_second_commit")

	commands.Checkout(secondCommit.Hash)

	file1, _ := os.Stat("file_first_commit")
	file2, _ := os.Stat("file_second_commit")

	if file1 == nil || file2 == nil {
		t.Fatal("one of the files was not reverted")
	}
}
