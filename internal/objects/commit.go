package objects

import (
	"fmt"
	"strings"
)

type Commit struct {
	Hash    string
	Tree    string
	Parents []string
	Author  string
	Message string
}

func (commit Commit) Stringify() string {
	var data string

	data += "tree " + commit.Tree + "\n"

	for _, parent := range commit.Parents {
		data += "parent " + parent + "\n"
	}

	data += "author " + commit.Author + "\n\n"

	data += commit.Message + "\n"

	return data
}

func (commit Commit) IsEmpty() bool {
	return commit.Hash == "" || commit.Tree == ""
}

func ParseCommit(hash string, data string) *Commit {
	lines := strings.Split(data, "\n")
	commit := Commit{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) < 2 {
			commit.Message += fmt.Sprintf("%s\n", line)
			continue
		}

		key, value := parts[0], parts[1]

		if key == "tree" {
			commit.Tree = value
		} else if key == "parent" {
			commit.Parents = append(commit.Parents, value)
		} else if key == "author" {
			commit.Author = value
		} else {
			commit.Message += line
		}
	}

	commit.Hash = hash

	return &commit
}
