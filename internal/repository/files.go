package repository

import "mgit/internal/objects"

func (repo *Repository) TrackedFiles(hash string) []string {
	commit := objects.ParseCommit(hash, repo.CatFile(hash))

	if commit.Tree == "" {
		return []string{}
	}

	tree := objects.ParseTree(commit.Tree, repo.CatFile(commit.Tree))

	var paths []string = make([]string, len(tree.Entries))

	for i, entry := range tree.Entries {
		paths[i] = entry.Path
	}

	return paths
}
