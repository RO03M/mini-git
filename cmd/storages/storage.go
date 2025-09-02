package storages

import (
	"fmt"
	"log"
	"mgit/cmd/paths"
	"os"
)

func SaveToStorage(hash string, content []byte) {
	hashDirName, hashFileName := hash[:2], hash[2:]

	objectDir := fmt.Sprintf("%s/objects/%s", paths.RepoName, hashDirName)
	os.MkdirAll(objectDir, 0755)

	os.WriteFile(fmt.Sprintf("%s/%s", objectDir, hashFileName), content, 0755)
}

func ReadFromStorage(hash string) []byte {
	hashDirName, hashFileName := hash[:2], hash[2:]
	objectPath := fmt.Sprintf("%s/objects/%s/%s", paths.RepoName, hashDirName, hashFileName)

	if _, err := os.Stat(objectPath); err != nil {
		log.Fatal(err, "r1")
	}

	file, err := os.ReadFile(objectPath)

	if err != nil {
		log.Fatal(err, "r2")
	}

	return file
}
