package repository

func (repo *Repository) TrackedFiles(hash string) []string {
	commit := repo.GetCommit(hash)

	if commit.Tree == "" {
		return []string{}
	}

	tree := repo.GetTree(commit.Tree)

	var paths []string = make([]string, len(tree.Entries))

	for i, entry := range tree.Entries {
		paths[i] = entry.Path
	}

	return paths
}
