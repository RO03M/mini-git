package commands

import (
	"fmt"
	"mgit/cmd/structures/commit"
)

func Tracked() {
	paths := commit.HeadTrackedFilesTree()

	for _, path := range paths {
		fmt.Println(path)
	}

	fmt.Printf("\nTotal of %v tracked files\n", len(paths))
}
