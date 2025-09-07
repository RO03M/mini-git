package storage

import (
	"fmt"
	"log"
	"mgit/cmd/paths"
	"os"
)

func objectsPath(hash string) string {
	hashDirName, hashFileName := hash[:2], hash[2:]
	return fmt.Sprintf("%s/objects/%s/%s", paths.RepoName, hashDirName, hashFileName)
}

func GetObjectByHash(hash string) []byte {
	hashDirName, hashFileName := hash[:2], hash[2:]
	objectPath := fmt.Sprintf("%s/objects/%s/%s", paths.RepoName, hashDirName, hashFileName)

	if _, err := os.Stat(objectPath); err != nil {
		log.Fatal(err, "r1")
	}

	file, err := os.ReadFile(objectPath)

	if err != nil {
		log.Fatal(err, "r2")
	}

	decompressedFile := Decompress(file)

	return decompressedFile
}

func Exists(hash string) bool {
	path := objectsPath(hash)

	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
