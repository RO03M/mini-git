package diff_test

import (
	"mgit/internal/diff"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFileDiff(t *testing.T) {
	a := strings.Split("line1\nline2\nline3\n", "\n")
	b := strings.Split("line1\nline2 changed\nline3\nline4\n", "\n")

	diffs := diff.Diff(a, b)

	diff.PrintLineDiffs(diffs)
}

func TestSwapLines(t *testing.T) {
	old := []string{"A", "B", "C", "D"}
	new := []string{"A", "B", "X", "C", "D"}

	diffs := diff.Diff(old, new)

	diff.PrintLineDiffs(diffs)
}

func TestDiffRealFiles(t *testing.T) {
	aFile, _ := os.ReadFile("./data/A")
	bFile, _ := os.ReadFile("./data/B")

	a := strings.Split(string(aFile), "\n")
	b := strings.Split(string(bFile), "\n")

	diffs := diff.Diff(a, b)

	diff.PrintLineDiffs(diffs)
}

func TestDiffEqualLines(t *testing.T) {
	old := []string{"a", "b", "c"}
	new := []string{"a", "b", "c"}

	got := diff.Diff(old, new)
	want := []diff.LineDiff{
		{OldContent: "a", NewContent: "a", Type: diff.DiffEqual, Line: 1},
		{OldContent: "b", NewContent: "b", Type: diff.DiffEqual, Line: 2},
		{OldContent: "c", NewContent: "c", Type: diff.DiffEqual, Line: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Equal lines diff mismatch:\n got  %+v\n want %+v", got, want)
	}
}

func TestDiffInsertLines(t *testing.T) {
	old := []string{"a"}
	new := []string{"a", "b", "c"}

	got := diff.Diff(old, new)
	want := []diff.LineDiff{
		{OldContent: "a", NewContent: "a", Type: diff.DiffEqual, Line: 1},
		{OldContent: "", NewContent: "b", Type: diff.DiffInsert, Line: 2},
		{OldContent: "", NewContent: "c", Type: diff.DiffInsert, Line: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert lines diff mismatch:\n got  %+v\n want %+v", got, want)
	}
}

func TestDiffDeleteLines(t *testing.T) {
	old := []string{"a", "b", "c"}
	new := []string{"a"}

	got := diff.Diff(old, new)
	want := []diff.LineDiff{
		{OldContent: "a", NewContent: "a", Type: diff.DiffEqual, Line: 1},
		{OldContent: "b", NewContent: "", Type: diff.DiffDelete, Line: 2},
		{OldContent: "c", NewContent: "", Type: diff.DiffDelete, Line: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Delete lines diff mismatch:\n got  %+v\n want %+v", got, want)
	}
}

func TestDiffModifiedLines(t *testing.T) {
	old := []string{"a", "b", "c"}
	new := []string{"a", "x", "y"}

	got := diff.Diff(old, new)
	want := []diff.LineDiff{
		{OldContent: "a", NewContent: "a", Type: diff.DiffEqual, Line: 1},
		{OldContent: "b", NewContent: "x", Type: diff.DiffModified, Line: 2},
		{OldContent: "c", NewContent: "y", Type: diff.DiffModified, Line: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Modified lines diff mismatch:\n got  %+v\n want %+v", got, want)
	}
}
