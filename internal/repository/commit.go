package repository

import (
	"log"
	"maps"
	"mgit/internal/objects"
	"slices"
)

func (repo *Repository) LastCommit() *objects.Commit {
	head := repo.RevParse("HEAD")

	if head == "" {
		return nil
	}

	object, err := repo.storage.Get(head)

	if err != nil {
		return nil
	}

	commit := objects.ParseCommit(head, string(object))

	if commit.IsEmpty() {
		return nil
	}

	return commit
}

func (repo *Repository) treeEntriesFromIndex() []objects.TreeEntry {
	items := repo.index.Items

	var entries []objects.TreeEntry = make([]objects.TreeEntry, len(items))

	for i, item := range slices.Collect(maps.Values(items)) {
		entries[i] = objects.TreeEntry{
			Type: objects.EntryTypeBlob,
			Hash: item.Hash,
			Path: item.Path,
		}
	}

	return entries
}

func (repo *Repository) buildTree(fromCommit *objects.Commit) *objects.Tree {
	additions := repo.index.Additions()

	var entries []objects.TreeEntry = make([]objects.TreeEntry, len(additions))

	for i, item := range additions {
		entries[i] = objects.TreeEntry{
			Type: objects.EntryTypeBlob,
			Hash: item.Hash,
			Path: item.Path,
		}
	}

	tree := objects.Tree{
		Entries: entries,
	}

	if fromCommit != nil {
		lastTree := objects.ParseTree(fromCommit.Tree, repo.CatFile(fromCommit.Tree))

		if lastTree != nil {
			tree.Merge(lastTree)
		}
	}

	deletions := repo.index.Deletions()
	var deletionPaths []string = make([]string, len(deletions))

	for i, item := range deletions {
		deletionPaths[i] = item.Path
	}

	tree.RemoveEntries(deletionPaths)

	return &tree
}

func (repo *Repository) Commit(message string) *objects.Commit {
	if len(repo.index.Items) == 0 {
		log.Fatal("no staged files")
	}

	// tree := objects.Tree{
	// 	Entries: repo.treeEntriesFromIndex(),
	// }

	var parent string
	lastCommit := repo.LastCommit()

	if lastCommit != nil {
		parent = lastCommit.Hash
		// lastTree := objects.ParseTree(lastCommit.Tree, repo.CatFile(lastCommit.Tree))

		// if lastTree != nil {
		// 	tree.Merge(lastTree)
		// }
	}

	tree := repo.buildTree(lastCommit)

	treeHash, _ := repo.storage.Create([]byte(tree.Stringify()))
	tree.Hash = treeHash

	commit := objects.Commit{
		Tree:    treeHash,
		Parents: []string{parent},
		Author:  "Gabriel Romera",
		Message: message,
	}

	hash, err := repo.storage.Create([]byte(commit.Stringify()))

	if err != nil {
		log.Fatalf("failed to create commit: %v", err)
	}

	commit.Hash = hash
	repo.UpdateHeadPointer(commit.Hash)
	repo.index.Clear()

	return &commit
}
