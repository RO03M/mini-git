package paths

import (
	"os"
	"path/filepath"
)

func GetDirTree(dir string) []string {
	ignore := LoadIgnoreFile()

	entries, _ := os.ReadDir(dir)
	var paths []string

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())
		if ignore.Match(entryPath) {
			continue
		}
		if entry.IsDir() {
			paths = append(paths, GetDirTree(entryPath)...)
		} else {
			paths = append(paths, entryPath)
		}
	}

	return paths
}
