package branch

import (
	"fmt"
	"log"
	"mgit/cmd/paths"
	"mgit/cmd/storage"
	"os"
	"path"
	"path/filepath"
)

type Branch struct {
	Name string
	Hash string
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

func FindBranch(ref string) *Branch {
	root := storage.GetRoot()
	branchPath := root + "/refs/heads/" + path.Base(ref)
	file, err := os.ReadFile(branchPath)

	if err != nil {
		return nil
	}

	fmt.Println(filepath.Rel(root, branchPath))

	return &Branch{
		Name: ref,
		Hash: string(file),
	}
}

func NameToRef(name string) string {
	return fmt.Sprintf("refs/heads/%s", name)
}

func (branch *Branch) Ref() string {
	return fmt.Sprintf("refs/heads/%s", branch.Name)
}
