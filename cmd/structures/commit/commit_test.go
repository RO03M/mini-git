package commit_test

import (
	"fmt"
	"log"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/tree"
	"reflect"
	"strings"
	"testing"
)

func TestStringify(t *testing.T) {
	tree := tree.Tree{
		Hash:  "hash",
		Blobs: []blob.Blob{},
	}
	commit := commit.CreateCommit("commit message", "parent hash", &tree)
	commit.Author = "Romera"
	commit.Committer = "Romera"

	text := commit.Stringify()

	var sb strings.Builder
	sb.WriteString("commit " + commit.Hash + "\n")
	sb.WriteString("tree " + tree.Hash + "\n")
	sb.WriteString("parent parent hash\n")
	sb.WriteString("author Romera\n")
	sb.WriteString("committer Romera\n")
	sb.WriteString("\n")
	sb.WriteString("commit message\n")

	if text != sb.String() {
		log.Fatalf("Stringify is not what we expected\n%s%s", text, sb.String())
	}
}

func TestParse(t *testing.T) {
	c := commit.Commit{
		Hash:   "hash",
		Parent: "hash",
		Tree: &tree.Tree{
			Hash: "hash",
		},
		Author:    "Romera",
		Committer: "Romera",
		Message:   "commit message",
	}

	var sb strings.Builder
	sb.WriteString("commit hash\n")
	sb.WriteString("tree hash\n")
	sb.WriteString("parent hash\n")
	sb.WriteString("author Romera\n")
	sb.WriteString("committer Romera\n")
	sb.WriteString("\n")
	sb.WriteString("commit message\n")

	parsedCommit, _ := commit.Parse(sb.String())

	if !reflect.DeepEqual(c, *parsedCommit) {
		fmt.Println(c.Tree, parsedCommit.Tree)
		log.Fatalf("\ncommit mismatch: \ngot %+v\nwant %+v", *parsedCommit, c)
	}
	fmt.Println(c, *parsedCommit)
}
