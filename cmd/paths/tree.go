package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDirTree(dir string) []string {
	ignore := LoadIgnoreFile()
	fmt.Println(os.Getwd())
	entries, _ := os.ReadDir(dir)
	fmt.Println(dir, entries)
	var paths []string

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())
		fmt.Println(entryPath, ignore.Match(entry.Name()), entry.Name())
		if ignore.Match(entry.Name()) {
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
