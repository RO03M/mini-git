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
	CurrentBlob blob.Blob
	TargetBlob  blob.Blob
	Type        Operation
}

func blobs2Map(blobs []blob.Blob) map[string]string {
	var blobMap map[string]string = map[string]string{}

	for _, blob := range blobs {
		blobMap[blob.FilePath] = blob.Hash
	}

	return blobMap
}

func (current *Tree) Diff(target Tree) []TreeDiff {
	if len(current.Blobs) == 0 {
		current.LoadBlobs()
	}

	if len(target.Blobs) == 0 {
		target.LoadBlobs()
	}

	var diffs []TreeDiff = []TreeDiff{}
	currentBlobMap := blobs2Map(current.Blobs)
	targetBlobMap := blobs2Map(target.Blobs)

	for _, blobObj := range current.Blobs {
		if blobMatchHash, found := targetBlobMap[blobObj.FilePath]; found {
			if blobMatchHash == blobObj.Hash {
				diffs = append(diffs, TreeDiff{
					CurrentBlob: blobObj,
					TargetBlob: blob.Blob{
						Hash:     blobMatchHash,
						FilePath: blobObj.FilePath,
					},
					Type: DiffEqual,
				})
			} else {
				diffs = append(diffs, TreeDiff{
					CurrentBlob: blobObj,
					TargetBlob: blob.Blob{
						Hash:     blobMatchHash,
						FilePath: blobObj.FilePath,
					},
					Type: DiffModified,
				})
			}
		} else {
			diffs = append(diffs, TreeDiff{
				CurrentBlob: blobObj,
				TargetBlob: blob.Blob{
					Hash:     "",
					FilePath: blobObj.FilePath,
				},
				Type: DiffDelete,
			})
		}
	}

	for _, blobObj := range target.Blobs {
		if _, found := currentBlobMap[blobObj.FilePath]; found {
			continue
		}

		diffs = append(diffs, TreeDiff{
			CurrentBlob: blob.Blob{},
			TargetBlob:  blobObj,
			Type:        DiffInsert,
		})
	}

	return diffs
}
