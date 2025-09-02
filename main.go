package main

import (
	"fmt"
	"log"
	"mgit/cmd"
	"mgit/cmd/paths"
	"mgit/cmd/storages"
	"mgit/cmd/structures"
	"os"
	"os/exec"
	"strings"
)

func initGit() {
	if info, _ := os.Stat(paths.RepoName); info != nil {
		fmt.Println("Already initialized")
		return
	}

	os.MkdirAll(paths.RepoName, 0755)
	os.MkdirAll(fmt.Sprintf("%s/refs/heads", paths.RepoName), 0755)
	os.MkdirAll(fmt.Sprintf("%s/refs/objects", paths.RepoName), 0755)

	os.WriteFile(fmt.Sprintf("%s/HEAD", paths.RepoName), []byte("ref: refs/heads/master"), 0644)
	os.WriteFile(fmt.Sprintf("%s/index", paths.RepoName), nil, 0644)
	os.WriteFile(fmt.Sprintf("%s/refs/heads/master", paths.RepoName), nil, 0644)
}

func addFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s doesn't exist\n", path)
		return
	}

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	blob := structures.CreateBlob(file)
	storages.SaveToStorage(blob.Hash, blob.Content)

	storages.AddStage(path, blob.Hash)
}

func commitStaged(message string) {
	stages := storages.GetStages()

	var parentHash string
	lastCommit := structures.GetCommitFromHead()

	if lastCommit != nil {
		parentHash = lastCommit.Hash
	}

	if len(stages) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	blobs := structures.StageObjectsToBlobs(stages)

	tree := structures.CreateTree(blobs)
	tree.Save()

	commit := structures.CreateCommit(message, parentHash, tree)
	commit.Save()

	storages.ClearStage()
	structures.UpdateHead(commit.Hash)

	fmt.Printf("Committed %v files\n\n", len(stages))
	fmt.Println(commit.Stringify())
}

func catFile(hash string) {
	object := storages.ReadFromStorage(hash)

	decompressed := cmd.Decompress(object)

	fmt.Println(string(decompressed))
}

func logCommits() {
	lastCommit := structures.GetCommitFromHead()

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

func checkout(ref string) {
	target := structures.CommitFromHash(ref)
	target.Tree.LoadBlobs()

	for _, blob := range target.Tree.Blobs {
		content := blob.ReadContent()
		os.WriteFile(blob.FilePath, content, 0644)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Custom git command made by Romera :)")
		fmt.Println("Run help to see the available commands")
		return
	}

	switch args[0] {
	case "init":
		initGit()
	case "add":
		addFile(args[1])
	case "commit":
		commitStaged(args[1])
	case "cat-file":
		catFile(args[1])
	case "log":
		logCommits()
	case "checkout":
		checkout(args[1])
	default:
		fmt.Println("Invalid command")
	}
}
