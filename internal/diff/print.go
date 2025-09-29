package diff

import (
	"fmt"
	"mgit/internal/plumbing"
)

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

func PrintLineDiffsColored(diffs []LineDiff) {
	for _, diff := range diffs {
		switch diff.Type {
		case DiffEqual, DiffMoved:
			fmt.Printf("  %s\n", diff.NewContent)
		case DiffInsert:
			plumbing.PrintfColor(plumbing.ColorGreen, "+ %s\n", diff.NewContent)
		case DiffDelete:
			plumbing.PrintfColor(plumbing.ColorRed, "- %s\n", diff.OldContent)
		case DiffModified:
			plumbing.PrintfColor(plumbing.ColorRed, "- %s\n", diff.OldContent)
			plumbing.PrintfColor(plumbing.ColorGreen, "+ %s\n", diff.NewContent)
		}
	}
}

func DiffsToText(diffs []LineDiff) string {
	text := ""

	for _, diff := range diffs {
		switch diff.Type {
		case DiffEqual, DiffMoved:
			text += fmt.Sprintf("  %s\n", diff.NewContent)
		case DiffInsert:
			text += fmt.Sprintf("+ %s\n", diff.NewContent)
		case DiffDelete:
			text += fmt.Sprintf("- %s\n", diff.OldContent)
		case DiffModified:
			text += fmt.Sprintf("- %s\n", diff.OldContent)
			text += fmt.Sprintf("+ %s\n", diff.NewContent)
		}
	}

	return text
}
