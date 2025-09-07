package storage

import (
	"fmt"
	"mgit/cmd/paths"
	"os"
)

func Init() {
	if info, _ := os.Stat(paths.RepoName); info != nil {
		fmt.Println("Already initialized")
		return
	}

	os.MkdirAll(paths.RepoName, 0755)
	os.MkdirAll(fmt.Sprintf("%s/refs/heads", paths.RepoName), 0755)
	os.MkdirAll(fmt.Sprintf("%s/refs/objects", paths.RepoName), 0755)

	os.WriteFile(fmt.Sprintf("%s/HEAD", paths.RepoName), []byte("ref: refs/heads/master"), 0644)
	os.WriteFile(fmt.Sprintf("%s/index", paths.RepoName), nil, 0644)
	os.WriteFile(fmt.Sprintf("%s/refs/heads/master", paths.RepoName), nil, 0644)
}
