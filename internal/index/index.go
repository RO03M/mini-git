package index

import (
	"log"
	"os"
	"strings"
)

type Index struct {
	Path  string
	Items map[string]Item
}

func Open(path string) *Index {
	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("failed to open index: %v", err)
	}

	lines := strings.Split(string(file), "\n")
	var items map[string]Item = map[string]Item{}

	for _, line := range lines {
		item := Parse(line)

		if item.Hash == "" {
			continue
		}

		items[item.Hash] = item
	}

	return &Index{
		Path:  path,
		Items: items,
	}
}

func (index *Index) Add(path string, hash string) {
	item := Item{
		Path:   path,
		Hash:   hash,
		Action: ActionAdd,
	}

	index.Items[path] = item
}

func (index *Index) AddRm(path string) {
	item := Item{
		Path:   path,
		Hash:   "",
		Action: ActionDelete,
	}

	index.Items[path] = item
}

func (index *Index) WriteBuffer() error {
	file, err := os.OpenFile(index.Path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	for _, item := range index.Items {
		file.WriteString(item.Stringify())
		file.WriteString("\n")
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return nil
}

func (index *Index) Clear() error {
	index.Items = map[string]Item{}
	return os.Truncate(index.Path, 0)
}

func (index *Index) Additions() []Item {
	additions := []Item{}

	for _, item := range index.Items {
		if item.Action != ActionAdd {
			continue
		}

		additions = append(additions, item)
	}

	return additions
}

func (index *Index) Deletions() []Item {
	deletions := []Item{}

	for _, item := range index.Items {
		if item.Action != ActionDelete {
			continue
		}

		deletions = append(deletions, item)
	}

	return deletions
}
