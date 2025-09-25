package objects

type Operation int8

const (
	DiffDelete   Operation = -1
	DiffEqual    Operation = 0
	DiffInsert   Operation = 1
	DiffModified Operation = 2
)

// Unir os paths
type TreeDiff struct {
	Path        string
	CurrentHash string
	TargetHash  string
	Type        Operation
}

func entryPathHash(entries []TreeEntry) map[string]string {
	var entryMap map[string]string = map[string]string{}

	for _, entry := range entries {
		entryMap[entry.Path] = entry.Hash
	}

	return entryMap
}

func (current Tree) Diff(target Tree) []TreeDiff {
	var diffs []TreeDiff = []TreeDiff{}
	currentEntryMap := entryPathHash(current.Entries)
	targetEntryMap := entryPathHash(target.Entries)

	for _, entry := range current.Entries {
		if targetHash, found := targetEntryMap[entry.Path]; found {
			if targetHash == entry.Hash {
				diffs = append(diffs, TreeDiff{
					CurrentHash: entry.Hash,
					Path:        entry.Path,
					TargetHash:  targetHash,
					Type:        DiffEqual,
				})
			} else {
				diffs = append(diffs, TreeDiff{
					CurrentHash: entry.Hash,
					Path:        entry.Path,
					TargetHash:  targetHash,
					Type:        DiffModified,
				})
			}
		} else {
			diffs = append(diffs, TreeDiff{
				CurrentHash: entry.Hash,
				Path:        entry.Path,
				TargetHash:  "",
				Type:        DiffDelete,
			})
		}
	}

	for _, entry := range target.Entries {
		if _, found := currentEntryMap[entry.Path]; found {
			continue
		}

		diffs = append(diffs, TreeDiff{
			Path:        entry.Path,
			CurrentHash: "",
			TargetHash:  entry.Hash,
			Type:        DiffInsert,
		})
	}

	return diffs
}
