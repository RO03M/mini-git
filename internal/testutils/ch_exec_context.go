package testutils

import (
	"log"
	"os"
	"testing"
)

func ChDirToTemp(t *testing.T) {
	tmpdir := t.TempDir()
	if err := os.Chdir(tmpdir); err != nil {
		log.Fatal(err)
	}
}
