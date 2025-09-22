package storage_test

import (
	"mgit/cmd/storage"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestGetCoreDirAbove(t *testing.T) {
	testutils.ChDirToTemp(t)
	storage.Init()
	os.MkdirAll("1/2/3/4", 0700)
	os.Chdir("1/2/3/4")

	mgitPath := storage.GetRoot()

	if mgitPath == "" {
		t.Fatal("no mgit path was found")
	}

	t.Log(mgitPath)
}

func TestIfDetectsNoMGitDir(t *testing.T) {
	testutils.ChDirToTemp(t)

	mgitPath := storage.GetRoot()

	if mgitPath != "" {
		t.Fatal("mgitPath should be empty")
	}
}
