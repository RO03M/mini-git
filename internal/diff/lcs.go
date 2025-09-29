package diff

import (
	"fmt"
	"slices"
)

func lcsMatrix(a []string, b []string) [][]int {
	var matrix [][]int = make([][]int, len(a)+1)

	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			charA := a[i-1]
			charB := b[j-1]

			if charA == charB {
				diagonal := matrix[i-1][j-1]
				matrix[i][j] = diagonal + 1
				continue
			}

			left := matrix[i][j-1]
			top := matrix[i-1][j]

			maxNeighbor := max(left, top)
			matrix[i][j] = maxNeighbor
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func Lcs(a []string, b []string) []string {
	matrix := lcsMatrix(a, b)

	elements := []string{}
	i := len(a)
	j := len(b)

	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			elements = append(elements, a[i-1])
			i--
			j--
		} else if matrix[i-1][j] > matrix[i][j-1] {
			i--
		} else {
			j--
		}
	}

	slices.Reverse(elements)

	return elements
}
