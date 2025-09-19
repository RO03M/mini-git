package tree

import "mgit/cmd/structures/blob"

type Operation int8

const (
	DiffDelete   Operation = -1
	DiffEqual    Operation = 0
	DiffInsert   Operation = 1
	DiffModified Operation = 2
)

type TreeDiff struct {
	Blob blob.Blob
	Type Operation
}

func blobs2Map(blobs []blob.Blob) map[string]string {
	var blobMap map[string]string = map[string]string{}

	for _, blob := range blobs {
		blobMap[blob.FilePath] = blob.Hash
	}

	return blobMap
}

func (tree *Tree) Diff(tree2 Tree) []TreeDiff {
	if len(tree.Blobs) == 0 {
		tree.LoadBlobs()
	}

	if len(tree2.Blobs) == 0 {
		tree2.LoadBlobs()
	}

	var diffs []TreeDiff = []TreeDiff{}
	blobMapTree1 := blobs2Map(tree.Blobs)
	blobMapTree2 := blobs2Map(tree2.Blobs)

	for _, blob := range tree.Blobs {
		if match, found := blobMapTree2[blob.FilePath]; found {
			if match == blob.Hash {
				diffs = append(diffs, TreeDiff{
					Blob: blob,
					Type: DiffEqual,
				})
			} else {
				diffs = append(diffs, TreeDiff{
					Blob: blob,
					Type: DiffModified,
				})
			}
		} else {
			diffs = append(diffs, TreeDiff{
				Blob: blob,
				Type: DiffDelete,
			})
		}
	}

	for _, blob := range tree2.Blobs {
		if _, found := blobMapTree1[blob.FilePath]; found {
			continue
		}

		diffs = append(diffs, TreeDiff{
			Blob: blob,
			Type: DiffInsert,
		})
	}

	return diffs
}
