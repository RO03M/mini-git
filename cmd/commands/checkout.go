package commands

import (
	"fmt"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
	"os"
)

func Checkout(ref string) {
	currentCommit := commit.GetCommitFromHead()

	target := commit.FromHash(ref)
	fmt.Println(currentCommit.Hash, target.Hash)
	target.Tree.LoadBlobs()

	diff1 := currentCommit.Tree.Diff(*target.Tree)
	// diff2 := target.Tree.Diff(*currentCommit.Tree)
	// fmt.Println(diff1)
	// fmt.Println(diff2)

	for _, diff := range diff1 {
		switch diff.Type {
		case tree.DiffDelete:
			fmt.Println(diff.Blob.FilePath, diff.Type, "delete")
			os.Remove(diff.Blob.FilePath)
		case tree.DiffModified, tree.DiffInsert, tree.DiffEqual:
			// O diff está correto, mas o blob de referência está vindo do currentCommit, e deveria vir do target
			// Isso pq o target é o que tem o conteúdo que desejamos colocar no arquivo
			content := diff.Blob.ReadContent()
			os.WriteFile(diff.Blob.FilePath, content, 0644)
			fmt.Println(diff.Blob.FilePath, diff.Type, "upsert", string(content))
		}
	}

	head.UpdateHead(target.Hash)
	// for _, blob := range target.Tree.Blobs {
	// 	content := blob.ReadContent()
	// 	os.WriteFile(blob.FilePath, content, 0644)
	// }
}
