package repository

import (
	"os"
	"path/filepath"
)

func (repo *Repository) Tracked(hash string) []string {
	commit := repo.GetCommit(hash)

	if commit == nil || commit.Tree == "" {
		return []string{}
	}

	tree := repo.GetTree(commit.Tree)

	var paths []string = make([]string, len(tree.Entries))

	for i, entry := range tree.Entries {
		paths[i] = entry.Path
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

	for _, path := range trackedPaths {
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
