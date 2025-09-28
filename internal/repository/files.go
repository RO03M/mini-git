package repository

import (
	"os"
	"path/filepath"
)

// returns a map of tracked file given a commit hash. key is the path, value is the corresponding hash
func (repo *Repository) Tracked(hash string) map[string]string {
	commit := repo.GetCommit(hash)

	if commit == nil || commit.Tree == "" {
		return map[string]string{}
	}

	tree := repo.GetTree(commit.Tree)

	paths := map[string]string{}

	for _, entry := range tree.Entries {
		paths[entry.Path] = entry.Hash
	}

	return paths
}

func (repo *Repository) Trackable(dir string) []string {
	entries, _ := os.ReadDir(dir)
	var paths []string

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())

		if repo.ignore.Match(entry.Name()) {
			continue
		}
		if entry.IsDir() {
			paths = append(paths, repo.Trackable(entryPath)...)
		} else {
			paths = append(paths, entryPath)
		}
	}

	return paths
}

func (repo *Repository) Untracked() []string {
	trackablePaths := repo.Trackable(".")
	trackedPaths := repo.Tracked(repo.RevParse("HEAD"))

	var trackedMap map[string]bool = make(map[string]bool)

	for path, _ := range trackedPaths {
		trackedMap[path] = true
	}

	for _, item := range repo.index.Items {
		trackedMap[item.Path] = true
	}

	var untrackedPaths []string = []string{}

	for _, path := range trackablePaths {
		if _, found := trackedMap[path]; found {
			continue
		}

		untrackedPaths = append(untrackedPaths, path)
	}

	return untrackedPaths
}
