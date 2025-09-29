package commands

import (
	"log"
	"mgit/internal/diff"
	"mgit/internal/plumbing"
	"mgit/internal/repository"
	"os"
	"strings"
)

func handleRefs(refA string, refB string) {
	repo := repository.Open()

	commitA := repo.GetCommit(repo.RevParse(refA))
	commitB := repo.GetCommit(repo.RevParse(refB))

	if commitA == nil {
		log.Fatalf("%s not found", refA)
	}

	if commitB == nil {
		log.Fatalf("%s not found", refB)
	}

	treeA := repo.GetTree(commitA.Tree)
	treeB := repo.GetTree(commitB.Tree)

	treeAEntryMap := map[string]string{}

	type ObjectDiff struct {
		Path  string
		Diffs []diff.LineDiff
	}

	objectsDiffs := []ObjectDiff{}

	for _, entry := range treeA.Entries {
		content, _ := repo.CatFile(entry.Hash)
		treeAEntryMap[entry.Path] = content
	}

	for _, entry := range treeB.Entries {
		content, _ := repo.CatFile(entry.Hash)
		if entryA, found := treeAEntryMap[entry.Path]; found {
			entryDiffs := diff.Diff(strings.Split(content, "\n"), strings.Split(entryA, "\n"))
			objectsDiffs = append(objectsDiffs, ObjectDiff{
				Path:  entry.Path,
				Diffs: entryDiffs,
			})

			delete(treeAEntryMap, entry.Path)

			continue
		}

		entryDiffs := diff.Diff(strings.Split(content, "\n"), []string{})
		objectsDiffs = append(objectsDiffs, ObjectDiff{
			Path:  entry.Path,
			Diffs: entryDiffs,
		})
	}

	for path, content := range treeAEntryMap {
		entryDiffs := diff.Diff([]string{}, strings.Split(content, "\n"))

		objectsDiffs = append(objectsDiffs, ObjectDiff{
			Path:  path,
			Diffs: entryDiffs,
		})
	}

	for _, objectDiffs := range objectsDiffs {
		plumbing.PrintfColor(plumbing.ColorCyan, "@file %s\n", objectDiffs.Path)
		diff.PrintLineDiffsColored(objectDiffs.Diffs)
	}
}

func handleFiles(refA string, refB string) {
	fileA, _ := os.ReadFile(refA)
	fileB, _ := os.ReadFile(refB)

	diffs := diff.Diff(strings.Split(string(fileA), "\n"), strings.Split(string(fileB), "\n"))

	diff.PrintLineDiffsColored(diffs)
}

func Diff(args ...string) {
	if len(args) != 2 {
		return
	}

	refA, refB := args[0], args[1]

	statA, _ := os.Stat(refA)
	statB, _ := os.Stat(refB)

	if statA == nil || statB == nil {
		handleRefs(refA, refB)
		return
	}

	handleFiles(refA, refB)
}
