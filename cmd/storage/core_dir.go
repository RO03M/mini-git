package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

func findMGitRecursive(dir string) string {
	absolutePath, _ := filepath.Abs(dir)

	path := fmt.Sprintf("%s/.mgit", absolutePath)

	if stat, _ := os.Stat(path); stat != nil {
		return path
	}

	if absolutePath == "/" {
		return "" // break out of the recursive loop
	}

	return findMGitRecursive(filepath.Dir(absolutePath))
}

func GetRoot() string {
	return findMGitRecursive(".")
}
