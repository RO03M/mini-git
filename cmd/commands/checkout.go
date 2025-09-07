package commands

import (
	"mgit/cmd/structures/commit"
	"os"
)

func Checkout(ref string) {
	target := commit.FromHash(ref)
	target.Tree.LoadBlobs()

	for _, blob := range target.Tree.Blobs {
		content := blob.ReadContent()
		os.WriteFile(blob.FilePath, content, 0644)
	}
}
