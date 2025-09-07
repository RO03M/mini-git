package storage

import (
	"bufio"
	"fmt"
	"log"
	"mgit/cmd/paths"
	"os"
	"strings"
)

type StageObject struct {
	Path string
	Hash string
}

func Parse(text string) *StageObject {
	var parts []string = make([]string, 2)
	parts = strings.Split(text, " ")

	path, hash := parts[0], parts[1]

	return &StageObject{
		Path: path,
		Hash: hash,
	}
}

func Stringify(stageObject StageObject) []byte {
	return []byte(fmt.Sprintf("%s %s", stageObject.Path, stageObject.Hash))
}

func OpenStageFile() *os.File {
	path := fmt.Sprintf("%s/index", paths.RepoName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.WriteFile(path, []byte{}, 0755)
	}

	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func AddStage(path string, hash string) {
	stageFile := OpenStageFile()

	newRow := fmt.Sprintf("%s %s\n", path, hash)

	stageObject := FindStageFromPath(path)
	if stageObject != nil {
		return
	}

	_, err := stageFile.WriteString(newRow)

	if err != nil {
		log.Fatalf("Failed to add %s to stage %v", path, err)
	}

	stageFile.Close()
}

func ClearStage() {
	path := fmt.Sprintf("%s/index", paths.RepoName)
	err := os.Truncate(path, 0)

	if err != nil {
		log.Fatal(err, "smcs")
	}
}

func GetStages() []StageObject {
	stageFile := OpenStageFile()
	stageScanner := bufio.NewScanner(stageFile)

	var stageObjects []StageObject = []StageObject{}

	for stageScanner.Scan() {
		stageObject := Parse(stageScanner.Text())
		stageObjects = append(stageObjects, *stageObject)
	}

	return stageObjects
}

func FindStageFromPath(targetPath string) *StageObject {
	stageFile := OpenStageFile()
	stageScanner := bufio.NewScanner(stageFile)

	for stageScanner.Scan() {
		stageObject := Parse(stageScanner.Text())
		if stageObject.Path == targetPath {
			return stageObject
		}
	}

	return nil
}
