package index

import (
	"fmt"
	"strings"
)

type StageAction string

const (
	ActionAdd    StageAction = "add"
	ActionDelete StageAction = "delete"
)

type Item struct {
	Path   string
	Hash   string
	Action StageAction
}

func Parse(text string) Item {
	parts := strings.Split(text, " ")

	if len(parts) != 3 {
		return Item{}
	}

	action, hash, path := parts[0], parts[1], parts[2]

	return Item{
		Path:   path,
		Hash:   hash,
		Action: StageAction(action),
	}
}

func (item *Item) Stringify() string {
	return fmt.Sprintf("%s %s %s", item.Action, item.Hash, item.Path)
}
