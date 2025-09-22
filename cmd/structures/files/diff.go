package files

import (
	"fmt"
	"math"
)

type Operation int8

const (
	DiffDelete   Operation = -1
	DiffEqual    Operation = 0
	DiffInsert   Operation = 1
	DiffModified Operation = 2
)

type LineDiff struct {
	OldContent string
	NewContent string
	Type       Operation
	Line       int
}

func Diff(old []string, new []string) []LineDiff {
	maxSize := int(math.Max(float64(len(old)), float64(len(new))))

	var diffs []LineDiff = make([]LineDiff, maxSize)

	line := 0

	for line < maxSize {
		var diffType Operation
		oldLine := ""
		newLine := ""

		if line < len(old) {
			oldLine = old[line]
		}

		if line < len(new) {
			newLine = new[line]
		}

		if oldLine == newLine {
			diffType = DiffEqual
		} else if oldLine == "" {
			diffType = DiffInsert
		} else if newLine == "" {
			diffType = DiffDelete
		} else {
			diffType = DiffModified
		}

		diffs[line] = LineDiff{
			OldContent: oldLine,
			NewContent: newLine,
			Type:       diffType,
			Line:       line + 1,
		}

		line++
	}

	return diffs
}

func PrintLineDiffs(diffs []LineDiff) {
	for _, diff := range diffs {
		switch diff.Type {
		case DiffEqual:
			fmt.Printf("%v %s\n", diff.Line, diff.NewContent)
		case DiffInsert:
			fmt.Printf("%v + %s\n", diff.Line, diff.NewContent)
		case DiffDelete:
			fmt.Printf("%v - %s\n", diff.Line, diff.OldContent)
		case DiffModified:
			fmt.Printf("%v - %s\n", diff.Line, diff.OldContent)
			fmt.Printf("%v + %s\n", diff.Line, diff.NewContent)
		}
	}
}
