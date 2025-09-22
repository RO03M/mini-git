package paths

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// var ignore *ignoreFile

var DEFAULT_IGNORE = []string{".mgit", ".git"}

type ignoreFile struct {
	raw      []byte
	lines    []string
	patterns []*regexp.Regexp
}

func LoadIgnoreFile() *ignoreFile {
	// if ignore != nil {
	// 	return ignore
	// }

	file, _ := os.ReadFile(".gitignore")

	ignore := &ignoreFile{
		raw:   file,
		lines: strings.Split(string(file), "\n"),
	}

	ignore.lines = append(ignore.lines, DEFAULT_IGNORE...)

	for _, line := range ignore.lines {
		rule, err := regexp.Compile("^" + line + "$")

		if err != nil {
			continue
		}

		ignore.patterns = append(ignore.patterns, rule)
	}

	return ignore
}

func (ig *ignoreFile) Match(path string) bool {
	unixPath := filepath.ToSlash(path)

	for _, r := range ig.patterns {
		if r.MatchString(unixPath) {
			return true
		}
	}

	return false
}
