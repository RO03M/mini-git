package commands

import (
	"fmt"
	"log"
	"mgit/internal/objects"
	"mgit/internal/plumbing"
	"mgit/internal/repository"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func unixLogs(commits []*objects.Commit) {
	logs := ""

	for _, commit := range commits {
		logs += plumbing.SprintfColor(plumbing.ColorYellow, "commit: %s\n", commit.Hash)
		logs += fmt.Sprintf("Author: %s\n", commit.Author)
		logs += fmt.Sprint(commit.Message)
		logs += "\n"
	}

	lineCount := len(strings.Split(logs, "\n"))

	if lineCount < 50 {
		fmt.Println(logs)
		return
	}

	cmd := exec.Command("less", "-R")

	cmd.Stdin = strings.NewReader(logs)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func defaultLogs(commits []*objects.Commit) {
	logs := ""

	for _, commit := range commits {
		logs += fmt.Sprintf("commit: %s\n", commit.Hash)
		logs += fmt.Sprintf("Author: %s\n", commit.Author)
		logs += fmt.Sprint(commit.Message)
		logs += "\n"
	}

	fmt.Println(logs)
}

func Log() {
	repo := repository.Open()

	head := repo.RevParse("HEAD")

	if head == "" {
		log.Fatal("no commits at HEAD")
	}

	commitHistory := repo.CommitHistory(head)

	switch runtime.GOOS {
	case "linux":
		unixLogs(commitHistory)
	default:
		defaultLogs(commitHistory)
	}
}
