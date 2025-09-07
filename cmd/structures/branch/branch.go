package branch

import (
	"log"
	"mgit/cmd/paths"
	"os"
)

type Branch struct {
	Name string
}

func LoadBranch(name string) *Branch {
	path := paths.GetBranchPath(name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
	}

	_, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return &Branch{
		Name: name,
	}
}
