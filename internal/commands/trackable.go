package commands

import (
	"fmt"
	"mgit/internal/repository"
	"path/filepath"
	"strings"
	"time"
)

func Trackable() {
	start := time.Now()
	repo := repository.Open()

	entries := repo.Trackable(filepath.Dir(repo.DotPath))

	fmt.Println(strings.Join(entries, "\n"))

	fmt.Printf("\nEnd of trackable files.\nTotal: %v\nTook: %s\n", len(entries), time.Since(start))
}
