package commit_test

import (
	"log"
	"mgit/cmd/commands"
	"mgit/cmd/structures/commit"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestRemovedFiles(t *testing.T) {
	testutils.ChDirToTemp(t)

	filename := "file"

	os.WriteFile(filename, []byte{}, 0644)

	commands.Init()
	commands.Add(filename)
	commands.Commit("first")

	os.Remove(filename)

	removedFiles := commit.GetRemovedFiles()

	if len(removedFiles) != 1 || removedFiles[0] != filename {
		log.Fatal("wrong removed files list returned")
	}
}
