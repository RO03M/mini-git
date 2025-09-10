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

func TestIgnoreFile(t *testing.T) {
	testutils.ChDirToTemp(t)

	os.WriteFile(".gitignore", []byte(".mgit\nignore"), 0644)
	os.WriteFile("include", []byte{}, 0644)
	os.WriteFile("ignore", []byte{}, 0644)

	filePaths := paths.GetDirTree(".")
	fmt.Println(filePaths)
	if len(filePaths) != 2 {
		t.Fatalf("wrong file paths size\nexpected: 2\ngot: %v", len(filePaths))
	}

	matchInclude := false
	hasIgnore := false

	for _, path := range filePaths {
		if path == "include" {
			matchInclude = true
			continue
		}

		if path == "ignore" {
			hasIgnore = true
			continue
		}
	}

	if !matchInclude {
		t.Fatal("include file is not in the list")
	}

	if hasIgnore {
		t.Fatal("ignore file is in the list")
	}
}
