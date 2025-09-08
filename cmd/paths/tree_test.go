package paths_test

import (
	"fmt"
	"mgit/cmd/paths"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestDirTree(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.WriteFile("file", []byte{}, 0644)
	os.MkdirAll("f1/f2/f3", 0700)
	os.WriteFile("f1/f2/f3/file", []byte{}, 0644)

	filePaths := paths.GetDirTree(".")

	fmt.Println(filePaths)

	if len(filePaths) != 2 {
		t.Fatalf("wrong file paths size\nexpected: 2\ngot: %v", len(filePaths))
	}
}
