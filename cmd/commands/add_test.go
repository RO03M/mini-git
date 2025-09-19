package commands_test

import (
	"mgit/cmd/commands"
	"mgit/cmd/stage"
	"mgit/cmd/storage"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestModifyAndStage(t *testing.T) {
	testutils.ChDirToTemp(t)
	commands.Init()

	os.WriteFile("file_to_be_staged", []byte("content"), 0644)

	commands.Add("./file_to_be_staged")

	manager := stage.Load()
	stageFiles := manager.AllObjects()

	if len(stageFiles) != 1 {
		t.Fatalf("Wrong size of staged files\nexpected: 1\ngot: %v", len(stageFiles))
	}

	if stageFiles[0].Path != "./file_to_be_staged" {
		t.Fatal("Wrong file path")
	}

	if stageFiles[0].Hash == "" {
		t.Fatal("No hash stored in the stage")
	}

	createdObject := storage.Exists(stageFiles[0].Hash)

	if !createdObject {
		t.Fatal("Object was not created for the staged file")
	}
}

func TestAddRemovedFile(t *testing.T) {
	testutils.ChDirToTemp(t)
	commands.Init()

	os.WriteFile("file_to_be_staged", []byte("content"), 0644)

	commands.Add("file_to_be_staged")

	commands.Commit("first")

	os.Remove("file_to_be_staged")

	commands.Add("file_to_be_staged")

	manager := stage.Load()

	if len(manager.AllObjects()) != 1 {
		t.Fatal("wrong staged files size")
	}
}
