package paths_test

import (
	"fmt"
	"mgit/cmd/paths"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestDirTree(t *testing.T) {
	fmt.Println(1)
	testutils.ChDirToTemp(t)

	os.WriteFile("file", []byte{}, 0644)
	os.MkdirAll("f1/f2/f3", 0700)
	os.WriteFile("f1/f2/f3/file", []byte{}, 0644)

	filePaths := paths.GetDirTree(".")

	if len(filePaths) != 2 {
		t.Fatalf("wrong file paths size\nexpected: 2\ngot: %v", len(filePaths))
	}
}

func TestIgnoreFile(t *testing.T) {
	fmt.Println(2)
	testutils.ChDirToTemp(t)

	os.WriteFile(".gitignore", []byte(".mgit\nignore"), 0644)
	os.WriteFile("include", []byte{}, 0644)
	os.WriteFile("ignore", []byte{}, 0644)

	fmt.Println(os.Getwd())
	filePaths := paths.GetDirTree(".")
	fmt.Println(os.Getwd())

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

func TestIgnoreDir(t *testing.T) {
	fmt.Println(3)
	testutils.ChDirToTemp(t)

	os.WriteFile(".gitignore", []byte(".mgit\nfile\ndirectory"), 0644)
	os.WriteFile("file", []byte{}, 0644)
	os.MkdirAll("directory", 0700)
	os.MkdirAll("sub", 0700)
	os.WriteFile("directory/subfile", []byte{}, 0644)
	os.WriteFile("sub/file", []byte{}, 0644) // shouldn't track because file is listed in the ignore file

	os.WriteFile("ok", []byte{}, 0644)
	os.MkdirAll("directoryok", 0700)
	os.WriteFile("directoryok/ok", []byte{}, 0644)

	filePaths := paths.GetDirTree(".")

	t.Log(filePaths)

	if len(filePaths) != 3 {
		t.Fatalf("wrong result length\nexpected: 3\ngot: %v", len(filePaths))
	}

	if filePaths[1] != "directoryok/ok" || filePaths[2] != "ok" {
		t.Fatal("wrong nodes tracked")
	}
}
