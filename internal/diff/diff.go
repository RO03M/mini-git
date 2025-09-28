package diff

type Operation int8

const (
	DiffDelete   Operation = -1
	DiffEqual    Operation = 0
	DiffInsert   Operation = 1
	DiffModified Operation = 2
	DiffMoved    Operation = 3
)

type LineDiff struct {
	OldContent string
	NewContent string
	Type       Operation
	Line       int
}

func Diff(old []string, new []string) []LineDiff {
	lcs := Lcs(old, new)

	diffs := []LineDiff{}

	line_old := 0
	line_new := 0
	for _, line := range lcs {
		for old[line_old] != line {
			diffs = append(diffs, LineDiff{
				OldContent: old[line_old],
				NewContent: line,
				Type:       DiffDelete,
				Line:       line_old,
			})
			line_old++
		}

		for new[line_new] != line {
			diffs = append(diffs, LineDiff{
				OldContent: line,
				NewContent: new[line_new],
				Type:       DiffInsert,
				Line:       line_new,
			})
			line_new++
		}

		diffs = append(diffs, LineDiff{
			Type:       DiffEqual,
			OldContent: old[line_old],
			NewContent: new[line_new],
			Line:       line_new,
		})

		line_new++
		line_old++
	}

	for line_old < len(old) {
		diffs = append(diffs, LineDiff{
			OldContent: old[line_old],
			NewContent: "",
			Type:       DiffDelete,
			Line:       line_old,
		})

		line_old++
	}

	for line_new < len(new) {
		diffs = append(diffs, LineDiff{
			OldContent: "",
			NewContent: new[line_new],
			Type:       DiffInsert,
			Line:       line_new,
		})

		line_new++
	}

	return diffs
}
