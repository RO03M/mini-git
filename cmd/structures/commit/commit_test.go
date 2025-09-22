package commit_test

import (
	"log"
	"mgit/cmd/structures/blob"
	"mgit/cmd/structures/commit"
	"mgit/cmd/structures/tree"
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
		Message:   "\ncommit message\n",
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

	if c.Hash != parsedCommit.Hash {
		log.Fatalf("\ncommit hash mismatch: \ngot %+v\nwant %+v", parsedCommit.Hash, c.Hash)
	}

	if c.Author != parsedCommit.Author {
		log.Fatalf("\ncommit author mismatch: \ngot %+v\nwant %+v", parsedCommit.Author, c.Author)
	}

	if c.Committer != parsedCommit.Committer {
		log.Fatalf("\ncommit committer mismatch: \ngot %+v\nwant %+v", parsedCommit.Committer, c.Committer)
	}

	if c.Message != parsedCommit.Message {
		log.Fatalf("\ncommit message mismatch: \ngot %+v\nwant %+v", parsedCommit.Message, c.Message)
	}

	if c.Parent != parsedCommit.Parent {
		log.Fatalf("\ncommit parent mismatch: \ngot %+v\nwant %+v", parsedCommit.Parent, c.Parent)
	}

	if parsedCommit.Tree == nil {
		log.Fatal("tree is <nil>")
	}

	if c.Tree.Hash != parsedCommit.Tree.Hash {
		log.Fatalf("\ncommit tree mismatch: \ngot %+v\nwant %+v", parsedCommit.Tree.Hash, c.Tree.Hash)
	}
}
