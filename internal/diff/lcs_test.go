package diff_test

import (
	"fmt"
	"mgit/internal/diff"
	"os"
	"strings"
	"testing"
)

func TestLcs(t *testing.T) {
	a := []string{"1", "4", "5", "6", "9", "10", "11"}
	b := []string{"6", "4", "5", "9", "11"}

	fmt.Println(diff.Lcs(b, a))

	aa := "abcdef"
	bb := "aacf"

	fmt.Println(diff.Lcs(strings.Split(bb, ""), strings.Split(aa, "")))
}

func TestLcsRealCase(t *testing.T) {
	aFile, _ := os.ReadFile("./data/A")
	bFile, _ := os.ReadFile("./data/B")

	a := strings.Split(string(aFile), "\n")
	b := strings.Split(string(bFile), "\n")

	for _, line := range diff.Lcs(a, b) {
		fmt.Println(line)
	}
}
