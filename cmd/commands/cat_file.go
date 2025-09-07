package commands

import (
	"fmt"
	"mgit/cmd/storage"
)

func CatFile(hash string) {
	object := storage.GetObjectByHash(hash)

	fmt.Println(string(object))
}
