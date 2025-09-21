package commands

import (
	"fmt"
	"mgit/cmd/paths"
	"mgit/cmd/structures/commit"
)

func trackedPaths() []string {
	lastCommit := commit.GetCommitFromHead()

	if lastCommit == nil {
		return []string{}
	}

	tree := lastCommit.Tree

	if tree == nil {
		return []string{}
	}

	tree.LoadBlobs()
	blobs := tree.Blobs
	var paths []string = make([]string, len(blobs))

	for i, blob := range blobs {
		paths[i] = blob.FilePath
	}

	return paths
}

func Untracked() []string {
	trackablePaths := paths.GetDirTree(".")
	trackedPaths := trackedPaths()

	var trackedMap map[string]bool = make(map[string]bool)

	for _, path := range trackedPaths {
		trackedMap[path] = false
	}

	var untrackedPaths []string = []string{}

	for _, path := range trackablePaths {
		if _, found := trackedMap[path]; found {
			continue
		}

		untrackedPaths = append(untrackedPaths, path)
		fmt.Println(path)
	}

	return untrackedPaths
}
