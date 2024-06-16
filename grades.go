package main

import (
	"fmt"
	"math/rand"
)

func dp(input [][]int) (int, int) {
	rows := make([]map[int]int, len(input))
	cols := make([]map[int]int, len(input[0]))
	max

	for i := 0; i < len(input); i++ {
		for c := 0; c < len(input[i]); c++ {
			rows[i][input[i][c]]++
			cols[c][input[i][c]]++
		}
	}
	for i := 1; i <= 5; i++ {

	}

	// Output the corresponding row and column numbers
	// fmt.Printf("%d %d\n", bestRow+1, bestCol+1)
	return bestRow + 1, bestCol + 1
}

// Utility function to calculate the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Utility function to calculate the maximum grade in a submatrix
func maxGradeInSubmatrix(matrix [][]int, row, col int) int {
	maxGrade := 0
	for i := 0; i < len(matrix); i++ {
		if i != row {
			for j := 0; j < len(matrix[0]); j++ {
				if j != col {
					if matrix[i][j] > maxGrade {
						maxGrade = matrix[i][j]
					}
				}
			}
		}
	}
	return maxGrade
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		// Read the size of the matrix
		var n, m int
		fmt.Scan(&n, &m)

		// Read the grades matrix
		grades := make([]string, n)
		matrix := make([][]int, n)
		for j := 0; j < n; j++ {
			matrix[j] = make([]int, m)
			fmt.Scan(&grades[j])
			for c := 0; c < m; c++ {
				matrix[j][c] = int(grades[j][c] - '0')
			}
		}

		// Find the row and column to remove
		rowToRemove, colToRemove := dp(matrix)

		// Output the row and column to remove
		fmt.Println(rowToRemove, colToRemove)
	}
}
