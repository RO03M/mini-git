package stage

import (
	"bufio"
	"fmt"
	"log"
	"mgit/cmd/paths"
	"os"
)

type StageManager struct {
	objects        []Object
	deletedObjects []Object
}

func Load() StageManager {
	path := fmt.Sprintf("%s/index", paths.RepoName)
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	manager := StageManager{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		object := Parse(scanner.Text())

		if object.Action == StageModified {
			manager.objects = append(manager.objects, object)
		} else if object.Action == StageDeleted {
			manager.deletedObjects = append(manager.deletedObjects, object)
		}
	}

	return manager
}

func Truncate() {
	path := fmt.Sprintf("%s/index", paths.RepoName)
	err := os.Truncate(path, 0)

	if err != nil {
		log.Fatal(err)
	}
}

func BlankManager() StageManager {
	return StageManager{}
}

// Writes to the index storage all the objects added to the stage manager
func (manager *StageManager) Write() {
	path := fmt.Sprintf("%s/index", paths.RepoName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to stage: %v", err)
	}

	for _, object := range manager.AllObjects() {
		file.WriteString(object.Stringify())
		file.WriteString("\n")
	}

	file.Close()
}

// Add a normal file to stage storage
func (manager *StageManager) Stage(path string, hash string) {
	object := CreateObject(path, hash, StageModified)

	manager.objects = append(manager.objects, object)
}

// Add a deleted file
func (manager *StageManager) StageDeleted(path string) {
	object := CreateObject(path, "", StageDeleted)

	manager.deletedObjects = append(manager.objects, object)
}

func (manager *StageManager) AllObjects() []Object {
	var allObjects []Object = []Object{}

	allObjects = append(allObjects, manager.objects...)
	allObjects = append(allObjects, manager.deletedObjects...)

	return allObjects
}

func (manager StageManager) Objects() []Object {
	return manager.objects
}

func (manager StageManager) DeletedObjects() []Object {
	return manager.deletedObjects
}

func (manager StageManager) HasStages() bool {
	return len(manager.objects) != 0 || len(manager.deletedObjects) != 0
}

// Removed all staged files
func (manager *StageManager) Clear() {

}
