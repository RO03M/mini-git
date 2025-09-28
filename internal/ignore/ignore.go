package ignore

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var DefaultIgnore = []string{".mgit", ".git"}

type Ignore struct {
	Path     string
	patterns []*regexp.Regexp
}

func Open(path string) *Ignore {
	file, _ := os.ReadFile(path)

	ignore := Ignore{
		Path:     path,
		patterns: []*regexp.Regexp{},
	}

	lines := strings.Split(string(file), "\n")
	lines = append(lines, DefaultIgnore...)

	for _, line := range lines {
		rule, err := regexp.Compile("^" + line + "$")

		if err != nil {
			continue
		}

		ignore.patterns = append(ignore.patterns, rule)
	}

	return &ignore
}

func (ig *Ignore) Match(path string) bool {
	unixPath := filepath.ToSlash(path)

	for _, r := range ig.patterns {
		if r.MatchString(unixPath) {
			return true
		}
	}

	return false
}
