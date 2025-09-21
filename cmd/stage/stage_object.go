package stage

import (
	"fmt"
	"strings"
)

type StageAction string

const (
	StageModified StageAction = "modified"
	StageDeleted  StageAction = "deleted"
)

type Object struct {
	Path   string
	Hash   string
	Action StageAction
}

func Parse(text string) Object {
	var parts []string = make([]string, 2)
	parts = strings.Split(text, " ")

	path, hash, action := parts[0], parts[1], parts[2]

	return CreateObject(path, hash, StageAction(action))
}

func CreateObject(path string, hash string, action StageAction) Object {
	return Object{
		Path:   path,
		Hash:   hash,
		Action: action,
	}
}

func (stageObject *Object) Stringify() string {
	return fmt.Sprintf("%s %s %s", stageObject.Path, stageObject.Hash, stageObject.Action)
}
