package objects

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type TreeEntryType string

const (
	EntryTypeBlob TreeEntryType = "blob"
	EntryTypeTree TreeEntryType = "tree"
)

type TreeEntry struct {
	Type TreeEntryType
	Hash string
	Path string
}

type Tree struct {
	Hash    string
	Entries []TreeEntry
}

func ParseTree(hash string, data string) *Tree {
	tree := Tree{
		Hash: hash,
	}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		if len(parts) != 3 {
			continue
		}

		entry := TreeEntry{
			Type: TreeEntryType(parts[0]),
			Hash: parts[1],
			Path: parts[2],
		}

		tree.Entries = append(tree.Entries, entry)
	}

	return &tree
}

func (tree *Tree) Stringify() string {
	var lines []string = make([]string, len(tree.Entries))

	for i, entry := range tree.Entries {
		lines[i] = fmt.Sprintf("%s %s %s", entry.Type, entry.Hash, entry.Path)
	}

	return strings.Join(lines, "\n")
}

// key is the tree entry path
func (tree *Tree) entryMap() map[string]TreeEntry {
	var entryMap map[string]TreeEntry = map[string]TreeEntry{}

	for _, entry := range tree.Entries {
		entryMap[entry.Path] = entry
	}

	return entryMap
}

func (tree *Tree) Merge(merge *Tree) {
	if merge == nil {
		return
	}

	if len(merge.Entries) == 0 {
		return
	}

	var entryMap map[string]TreeEntry = map[string]TreeEntry{}

	for _, entry := range merge.Entries {
		entryMap[entry.Path] = entry
	}

	for _, entry := range tree.Entries {
		entryMap[entry.Path] = entry
	}

	tree.Entries = slices.Collect(maps.Values(entryMap))
}

func (tree *Tree) RemoveEntries(paths []string) {
	entryMap := tree.entryMap()

	for _, path := range paths {
		delete(entryMap, path)
	}

	tree.Entries = slices.Collect(maps.Values(entryMap))
}
