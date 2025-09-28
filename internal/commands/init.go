package commands

import (
	"fmt"
	"mgit/internal/repository"
	"os"
)

func Init() {
	stat, _ := os.Stat(".mgit")
	existed := stat != nil
	repo := repository.Initialize(".")

	if existed {
		fmt.Printf("Reinitialized mgit repository in %s\n", repo.DotPath)
	} else {
		fmt.Printf("Initialized mgit repository in %s\n", repo.DotPath)
	}
}
