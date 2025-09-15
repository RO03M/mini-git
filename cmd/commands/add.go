package commands

import (
	"fmt"
	"log"
	"mgit/cmd/stage"
	"mgit/cmd/storage"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/commit"
	"mgit/cmd/utils"
	"os"
)

// Should call this only if the file doesn't exist
func addRemovedFile(path string) {
	tracked := commit.HeadTrackedFilesTree()

	trackedMap := utils.StringSliceMap(tracked)

	if _, found := trackedMap[path]; !found {
		fmt.Println(path + " doesn't exist")
		return
	}

	manager := stage.BlankManager()
	manager.StageDeleted(path)
	manager.Write()
	fmt.Println("Added removed file ", path)
}

func addSingle(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		addRemovedFile(path)
		return
	}

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	blob := blob.CreateBlob(file)
	storage.Create(blob.Hash, blob.Content)

	stageManager := stage.BlankManager()
	stageManager.Stage(path, blob.Hash)

	stageManager.Write()
}

func Add(paths ...string) {
	for _, path := range paths {
		addSingle(path)
	}
}
