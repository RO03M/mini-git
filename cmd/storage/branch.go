package storage

import (
	"fmt"
	"os"
)

func CreateBranch(name string, initValue string) error {
	mgitDir := GetRoot()

	return os.WriteFile(fmt.Sprintf("%s/refs/heads/%s", mgitDir, name), []byte(initValue), 0644)
}

func UpdateBranch(name, value string) {
	mgitDir := GetRoot()

	os.WriteFile(fmt.Sprintf("%s/.mgit/refs/heads/%s", mgitDir, name), []byte(value), 0644)
}

func ReadBranch(name string) string {
	mgitDir := GetRoot()

	branchPath := fmt.Sprintf("%s/.mgit/refs/heads/%s", mgitDir, name)

	content, _ := os.ReadFile(branchPath)

	return string(content)
}

func DeleteBranch(name string) error {
	mgitDir := GetRoot()

	branchPath := fmt.Sprintf("%s/.mgit/refs/heads/%s", mgitDir, name)

	return os.Remove(branchPath)
}
