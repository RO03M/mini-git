package tree

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"maps"
	"mgit/cmd/storage"
	"mgit/cmd/structures/blob"
	"mgit/cmd/utils"
	"slices"
	"strings"
)

type Tree struct {
	Hash  string
	Blobs []blob.Blob
}

func CreateTree(blobs []blob.Blob) *Tree {
	return &Tree{
		// Hash:  hex.EncodeToString(hash),
		Blobs: blobs,
	}
}

func CreateMergedTree(prevTree *Tree, blobs []blob.Blob) *Tree {
	if prevTree == nil {
		return CreateTree(blobs)
	}

	if len(prevTree.Blobs) == 0 {
		prevTree.LoadBlobs()
	}

	var blobMap map[string]blob.Blob = map[string]blob.Blob{}

	for _, blobObj := range prevTree.Blobs {
		blobMap[blobObj.Hash] = blobObj
	}

	for _, blobObj := range blobs {
		blobMap[blobObj.Hash] = blobObj
	}

	uniqueBlobs := slices.Collect(maps.Values(blobMap))

	return CreateTree(uniqueBlobs)

	// var totalSize int = len(prevTree.Blobs) + len(blobs)
	// var mergedBlobs []blob.Blob = make([]blob.Blob, totalSize)
	// var i int = 0

	// for _, blob := range prevTree.Blobs {
	// 	mergedBlobs[i] = blob
	// 	i++
	// }

	// for _, blob := range blobs {
	// 	mergedBlobs[i] = blob
	// 	i++
	// }

	// return CreateTree(mergedBlobs)
}

func (tree *Tree) Stringify() string {
	var lines []string = make([]string, len(tree.Blobs))

	for i, blob := range tree.Blobs {
		lines[i] = fmt.Sprintf("blob %s %s", blob.Hash, blob.FilePath)
	}

	return strings.Join(lines, "\n")
}

func (tree *Tree) GenerateHash() {
	var blobHashes []string = make([]string, len(tree.Blobs))
	for i, blob := range tree.Blobs {
		blobHashes[i] = blob.Hash
	}

	hasher := sha1.New()
	hasher.Write([]byte(strings.Join(blobHashes, "")))
	hash := hasher.Sum(nil)

	tree.Hash = hex.EncodeToString(hash)
}

func (tree *Tree) Save() {
	tree.GenerateHash()
	storage.Create(tree.Hash, []byte(tree.Stringify()))
}

func (tree *Tree) LoadBlobs() {
	treeObject := storage.GetObjectByHash(tree.Hash)

	lines := strings.Split(string(treeObject), "\n")

	var blobs []blob.Blob

	for _, line := range lines {
		blob := blob.ParseBlob(line)
		if blob == nil {
			continue
		}

		blobs = append(blobs, *blob)
	}

	tree.Blobs = blobs
}

func (tree *Tree) RemoveBlobsByPath(paths ...string) {
	currentBlobs := tree.Blobs
	var newBlobs []blob.Blob = []blob.Blob{}
	pathMap := utils.StringSliceMap(paths)

	for _, blob := range currentBlobs {
		if _, found := pathMap[blob.FilePath]; found {
			continue
		}

		newBlobs = append(newBlobs, blob)
	}

	tree.Blobs = newBlobs
}
