package storage_test

import (
	"mgit/cmd/storage"
	"mgit/internal/testutils"
	"testing"
)

func TestShouldReturnRawHash(t *testing.T) {
	hash := "randomstringsupposedtobehash"

	ref := storage.GetHashFromRef(hash)

	if hash != ref {
		t.Fatalf("wrong ref\nexpected: %s\ngot: %s", hash, ref)
	}
}

func TestWithBranch(t *testing.T) {
	testutils.ChDirToTemp(t)

	storage.Init()
	storage.CreateBranch("test-branch", "branchhash")
	ref := "refs/heads/test-branch"

	hash := storage.GetHashFromRef(ref)

	if hash != "branchhash" {
		t.Fatalf("wrong hash returned from the branch ref\nexpected: branchhash\ngot: %s", hash)
	}
}
