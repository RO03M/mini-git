package structures

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"mgit/cmd"
	"mgit/cmd/storages"
	"strings"
)

type Commit struct {
	Hash      string
	Parent    string
	Tree      *Tree
	Author    string
	Committer string
	Message   string
}

func (commit Commit) Stringify() string {
	return fmt.Sprintf(`commit %s
tree %s
parent %s
author %s
commiter %s

%s
`, commit.Hash, commit.Tree.Hash, commit.Parent, commit.Author, commit.Committer, commit.Message)
}

func ParseCommit(data string) (*Commit, error) {
	lines := strings.Split(data, "\n")
	commit := Commit{}
	fmt.Println(data)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) < 2 {
			commit.Message += fmt.Sprintf("%s\n", line)
			continue
		}

		key, value := parts[0], parts[1]

		switch key {
		case "commit":
			commit.Hash = value
		case "tree":
			commit.Tree = &Tree{Hash: value}
		case "parent":
			commit.Parent = value
		case "author":
			commit.Author = value
		case "commiter":
			commit.Committer = value
		default:
			commit.Message += line
		}
	}

	return &commit, nil
}

func CreateCommit(message string, parent string, tree *Tree) *Commit {
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
	head, err := GetHeadHash()

	if err != nil {
		log.Fatal(err)
	}

	if head == "" {
		return nil
	}

	return CommitFromHash(head)
}

func CommitFromHash(hash string) *Commit {
	object := storages.ReadFromStorage(hash)

	decompressed := cmd.Decompress(object)

	commit, _ := ParseCommit(string(decompressed))

	return commit
}

func (commit *Commit) Parents() []Commit {
	current := commit

	var parents []Commit = []Commit{}

	for current.Parent != "" {
		parent := CommitFromHash(current.Parent)

		if parent == nil {
			break
		}

		parents = append(parents, *parent)
		current = parent
	}

	return parents
}

func (commit *Commit) Save() {
	commitObject := cmd.Compress([]byte(commit.Stringify()))
	storages.SaveToStorage(commit.Hash, commitObject)
}
