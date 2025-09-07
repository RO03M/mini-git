package storage_test

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"mgit/cmd/paths"
	"mgit/cmd/storage"
	"mgit/internal/testutils"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	testutils.ChDirToTemp(t)
	content := "file content"
	hash := storage.GenerateSha1([]byte("hash"))

	storage.Init()
	storage.Create(hash, []byte(content))

	if !storage.Exists(hash) {
		log.Fatal("Couldn't find the object")
	}
}

func TestCreateAndRead(t *testing.T) {
	testutils.ChDirToTemp(t)
	content := "file content"
	hash := storage.GenerateSha1([]byte("hash"))

	storage.Init()
	storage.Create(hash, []byte(content))

	object := storage.GetObjectByHash(hash)

	if string(object) != content {
		log.Fatalf("object is not the same as the original save\n\nobject: %v\ncontent: %v\n", string(object), content)
	}
}

func TestCreateIsZipping(t *testing.T) {
	testutils.ChDirToTemp(t)
	content := "file content"
	hash := storage.GenerateSha1([]byte("hash"))

	storage.Init()
	storage.Create(hash, []byte(content))

	hashDirName, hashFileName := hash[:2], hash[2:]
	path := fmt.Sprintf("%s/objects/%s/%s", paths.RepoName, hashDirName, hashFileName)
	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	bReader := bytes.NewReader(file)
	r, err := zlib.NewReader(bReader)

	if err != nil {
		log.Fatalf("File is not being compressed as default\n%v", err)
	}

	r.Close()
}
