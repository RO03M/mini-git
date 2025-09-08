package commands

import (
	"fmt"
	"mgit/cmd/paths"
	"strings"
	"time"
)

func Trackable() {
	start := time.Now()
	entries := paths.GetDirTree(".")

	fmt.Println(strings.Join(entries, "\n"))

	fmt.Printf("End of trackable files.\nTotal: %v\nTook: %s\n", len(entries), time.Since(start))
}
