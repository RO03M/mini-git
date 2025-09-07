package commit

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"mgit/cmd/storage"
	"mgit/cmd/structures/head"
	"mgit/cmd/structures/tree"
	"strings"
)

type Commit struct {
	Hash      string
	Parent    string
	Tree      *tree.Tree
	Author    string
	Committer string
	Message   string
}

func Parse(data string) (*Commit, error) {
	lines := strings.Split(data, "\n")
	commit := Commit{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) < 2 {
			commit.Message += fmt.Sprintf("%s\n", line)
			continue
		}

		key, value := parts[0], parts[1]

		switch key {
		case "commit":
			if commit.Hash != "" {
				commit.Message += line
				break
			}
			commit.Hash = value
		case "tree":
			commit.Tree = &tree.Tree{Hash: value}
		case "parent":
			commit.Parent = value
		case "author":
			commit.Author = value
		case "committer":
			commit.Committer = value
		default:
			commit.Message += line
		}
	}

	return &commit, nil
}

func CreateCommit(message string, parent string, tree *tree.Tree) *Commit {
	hasher := sha1.New()
	hashMessage := fmt.Sprintf("%s %s %s", message, parent, tree.Hash)
	hasher.Write([]byte(hashMessage))
	hash := hasher.Sum(nil)

	return &Commit{
		Hash:    hex.EncodeToString(hash),
		Parent:  parent,
		Tree:    tree,
		Message: message,
	}
}

func GetCommitFromHead() *Commit {
	head, err := head.GetHeadHash()

	if err != nil {
		log.Fatal(err)
	}

	if head == "" {
		return nil
	}

	return FromHash(head)
}

func FromHash(hash string) *Commit {
	object := storage.GetObjectByHash(hash)

	commit, _ := Parse(string(object))

	return commit
}

func (commit *Commit) Parents() []Commit {
	current := commit

	var parents []Commit = []Commit{}

	for current.Parent != "" {
		parent := FromHash(current.Parent)

		if parent == nil {
			break
		}

		parents = append(parents, *parent)
		current = parent
	}

	return parents
}

func (commit Commit) Stringify() string {
	return fmt.Sprintf(`commit %s
tree %s
parent %s
author %s
committer %s

%s
`, commit.Hash, commit.Tree.Hash, commit.Parent, commit.Author, commit.Committer, commit.Message)
}

func (commit *Commit) Save() {
	// commitObject := cmd.Compress([]byte(commit.Stringify()))
	storage.Create(commit.Hash, []byte(commit.Stringify()))
}
