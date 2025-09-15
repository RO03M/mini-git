package commands

import (
	"fmt"
	"mgit/cmd/structures/commit"
)

func RemovedFiles() {
	for _, path := range commit.GetRemovedFiles() {
		fmt.Println(path)
	}
}
