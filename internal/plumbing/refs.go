package plumbing

import (
	"regexp"
	"strings"
)

func IsRefPath(ref string) bool {
	for _, c := range ref {
		if c == '/' {
			return true
		}
	}

	return false
}

func BranchFromRef(ref string) string {
	branchRefPattern := regexp.MustCompile("^refs/heads/[A-Za-z0-9._/-]+$")

	match := branchRefPattern.MatchString(ref)

	if match {
		parts := strings.Split(ref, "/")

		return parts[len(parts)-1]
	}

	return ""
}
