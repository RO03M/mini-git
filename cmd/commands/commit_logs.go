package commands

import (
	"log"
	"mgit/cmd/structures/commit"
	"os"
	"os/exec"
	"strings"
)

func Log() {
	lastCommit := commit.GetCommitFromHead()

	if lastCommit == nil {
		log.Fatal("No commits were ever made")
	}

	var logs []string = []string{lastCommit.Stringify()}

	parents := lastCommit.Parents()

	for _, parent := range parents {
		logs = append(logs, parent.Stringify())
	}

	cmd := exec.Command("less")

	cmd.Stdin = strings.NewReader(strings.Join(logs, "\n"))
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
