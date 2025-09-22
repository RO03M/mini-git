package files_test

import (
	"mgit/cmd/structures/files"
	"strings"
	"testing"
)

func TestFileDiff(t *testing.T) {
	a := strings.Split("line1\nline2\nline3\n", "\n")
	b := strings.Split("line1\nline2 changed\nline3\nline4\n", "\n")

	diffs := files.Diff(a, b)

	files.PrintLineDiffs(diffs)
}
