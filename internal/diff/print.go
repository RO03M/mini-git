package diff

import "fmt"

func PrintLineDiffs(diffs []LineDiff) {
	for _, diff := range diffs {
		switch diff.Type {
		case DiffEqual, DiffMoved:
			fmt.Printf("  %s\n", diff.NewContent)
		case DiffInsert:
			fmt.Printf("+ %s\n", diff.NewContent)
		case DiffDelete:
			fmt.Printf("- %s\n", diff.OldContent)
		case DiffModified:
			fmt.Printf("- %s\n", diff.OldContent)
			fmt.Printf("+ %s\n", diff.NewContent)
		}
	}
}
