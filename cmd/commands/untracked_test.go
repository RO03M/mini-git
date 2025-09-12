package commands_test

import (
	"fmt"
	"mgit/cmd/commands"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestUntrackedFiles(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.MkdirAll("dir", 0700)
	os.WriteFile("file", []byte{}, 0644)
	os.WriteFile("dir/file1", []byte{}, 0644)
	os.WriteFile("dir/file2", []byte{}, 0644)
	os.WriteFile("dir/file3", []byte{}, 0644)
	os.WriteFile("dir/file4", []byte{}, 0644)

	commands.Init()

	commands.Add("dir/file1")
	commands.Add("dir/file2")

	commands.Commit("first")

	commands.Untracked()
	fmt.Println()

	os.WriteFile("dir/file5", []byte{}, 0644)

	commands.Untracked()
	fmt.Println()

	commands.Add("file")
	commands.Add("dir/file3")
	commands.Add("dir/file4")
	commands.Add("dir/file5")

	commands.Commit("second")

	commands.Untracked()
}
