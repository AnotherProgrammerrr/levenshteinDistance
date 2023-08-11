package main

import (
	"fmt"
)

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func levenshteinDistance(x, y string) int {
	n, m := len(x), len(y)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	matrix := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		matrix[i] = make([]int, m+1)
		matrix[i][0] = i
	}

	for j := 0; j <= m; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := 0
			if x[i-1] != y[j-1] {
				cost = 1
			}
			matrix[i][j] = min(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost,
			)
		}
	}
	return matrix[n][m]
}

func similarity(x, y string) float64 {
	dist := levenshteinDistance(x, y)
	maxLen := len(x)
	if len(y) > maxLen {
		maxLen = len(y)
	}
	return 1.0 - float64(dist)/float64(maxLen)
}

func main() {
	x := "hello world"
	y := "hallo welt"
	score := similarity(x, y)
	fmt.Printf("Similarity between '%s' and '%s' is %.2f\n", x, y, score)
}
