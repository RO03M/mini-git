package storage

import (
	"fmt"
	"os"
	"path"
)

func isObjectPath(ref string) bool {
	for _, char := range ref {
		if char == '/' {
			return true
		}
	}

	return false
}

func GetHashFromRef(ref string) string {
	if !isObjectPath(ref) {
		return ref
	}

	mgitPath := GetRoot()
	fmt.Println(path.Join(mgitPath, ref))
	file, _ := os.ReadFile(path.Join(mgitPath, ref))

	return string(file)
}
