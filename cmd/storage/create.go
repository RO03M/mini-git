package storage

import (
	"fmt"
	"mgit/cmd/paths"
	"os"
)

func Create(hash string, content []byte) {
	compressedContent := Compress(content)

	hashDirName, hashFileName := hash[:2], hash[2:]

	objectDir := fmt.Sprintf("%s/objects/%s", paths.RepoName, hashDirName)
	os.MkdirAll(objectDir, 0755)

	os.WriteFile(fmt.Sprintf("%s/%s", objectDir, hashFileName), compressedContent, 0755)
}
