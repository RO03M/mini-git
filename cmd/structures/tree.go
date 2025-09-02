package structures

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"mgit/cmd"
	"mgit/cmd/storages"
	"strings"
)

type Tree struct {
	Hash  string
	Blobs []Blob
}

func CreateTree(blobs []Blob) *Tree {
	var blobHashes []string = make([]string, len(blobs))
	for i, blob := range blobs {
		blobHashes[i] = blob.Hash
	}

	hasher := sha1.New()
	hasher.Write([]byte(strings.Join(blobHashes, "")))
	hash := hasher.Sum(nil)

	return &Tree{
		Hash:  hex.EncodeToString(hash),
		Blobs: blobs,
	}
}

func (tree *Tree) Stringify() string {
	var lines []string = make([]string, len(tree.Blobs))

	for i, blob := range tree.Blobs {
		lines[i] = fmt.Sprintf("blob %s %s", blob.Hash, blob.FilePath)
	}

	return strings.Join(lines, "\n")
}

func (tree *Tree) Save() {
	compressed := cmd.Compress([]byte(tree.Stringify()))
	storages.SaveToStorage(tree.Hash, compressed)
}

func (tree *Tree) LoadBlobs() {
	foo := storages.ReadFromStorage(tree.Hash)
	decompressed := cmd.Decompress(foo)

	lines := strings.Split(string(decompressed), "\n")

	var blobs []Blob

	for _, line := range lines {
		blob := ParseBlob(line)
		if blob == nil {
			continue
		}

		blobs = append(blobs, *blob)
	}

	tree.Blobs = blobs
}
