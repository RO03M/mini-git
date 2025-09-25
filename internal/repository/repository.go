package repository

import (
	"fmt"
	"log"
	"maps"
	"mgit/internal/index"
	"mgit/internal/storage"
	"os"
	"path/filepath"
	"slices"
)

type Repository struct {
	DotPath  string
	headPath string
	storage  *storage.Storage
	index    *index.Index
}

const DefaultDirPerm = 0755
const DefaultFilePerm = 0644

func findDotPath() (string, bool) {
	currentPath, err := os.Getwd()

	if err != nil {
		return "", false
	}

	for {
		dotPath := fmt.Sprintf("%s/.mgit", currentPath)

		if file, _ := os.Stat(dotPath); file != nil {
			return dotPath, true
		}

		if currentPath == "/" {
			return "", false
		}

		currentPath = filepath.Dir(currentPath)
	}
}

func newRepository(dotpath string) *Repository {
	repo := Repository{
		DotPath:  dotpath,
		headPath: filepath.Join(dotpath, "HEAD"),
		storage: &storage.Storage{
			ObjectsPath: filepath.Join(dotpath, "objects"),
		},
		index: index.Open(filepath.Join(dotpath, "index")),
	}

	return &repo
}

func Open() *Repository {
	dotpath, found := findDotPath()

	if !found {
		log.Fatal("Couldn't find .mgit directory. Initialize it with \"mgit init\"")
	}

	return newRepository(dotpath)
}

func Initialize(path string) *Repository {
	abspath, _ := filepath.Abs(path)

	stat, err := os.Stat(abspath)

	if err != nil {
		log.Fatalf("couldn't init repository: %v", err)
	}

	if stat == nil {
		log.Fatal("invalid path, stat not found")
	}

	if !stat.IsDir() {
		log.Fatal("the path is a file, you must provide a path to a directory")
	}

	abspath = filepath.Join(abspath, ".mgit")

	os.Mkdir(abspath, DefaultDirPerm)
	os.MkdirAll(filepath.Join(abspath, "refs/objects"), DefaultDirPerm)
	os.MkdirAll(filepath.Join(abspath, "refs/heads"), DefaultDirPerm)

	os.WriteFile(filepath.Join(abspath, "HEAD"), []byte("refs/heads/master"), DefaultFilePerm)
	os.WriteFile(filepath.Join(abspath, "index"), nil, DefaultFilePerm)
	os.WriteFile(filepath.Join(abspath, "refs/heads/master"), nil, DefaultFilePerm)

	return newRepository(abspath)
}

func (repo *Repository) PathFromDot(path string) string {
	abspath, _ := filepath.Abs(path)

	relpath, _ := filepath.Rel(filepath.Dir(repo.DotPath), abspath)

	return relpath
}

type StatusBody struct {
	Staged []index.Item
}

func (repo *Repository) Status() StatusBody {
	return StatusBody{
		Staged: slices.Collect(maps.Values(repo.index.Items)),
	}
}

func (repo *Repository) CatFile(hash string) (string, error) {
	object, err := repo.storage.Get(hash)

	if err != nil {
		return "", err
	}

	return string(object), nil
}
