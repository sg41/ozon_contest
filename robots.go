package main

import (
	"fmt"
	"math"
)

func findPathStart(result []string, robot [2]int, marker string) {
	for r, c := robot[0], robot[1]; r > 0 || c > 0; {
		if r > 0 && result[r-1][c] == '.' {
			result[r-1] = result[r-1][:c] + marker + result[r-1][c+1:]
			r--
		} else if c > 0 && result[r][c-1] == '.' {
			result[r] = result[r][:c-1] + marker + result[r][c:]
			c--
		} else {
			break
		}
	}
}
func findPathEnd(result []string, robot [2]int, marker string) {
	for r, c := robot[0], robot[1]; r < len(result) || c < len(result[0]); {
		if r < len(result)-1 && result[r+1][c] == '.' {
			result[r+1] = result[r+1][:c] + marker + result[r+1][c+1:]
			r++
		} else if c < len(result[0])-1 && result[r][c+1] == '.' {
			result[r] = result[r][:c+1] + marker + result[r][c+2:]
			c++
		} else {
			break
		}
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		A := [2]int{-1, -1}
		B := [2]int{-1, -1}

		// Чтение описания склада
		warehouse := make([]string, n)
		for j := 0; j < n; j++ {
			// warehouse[j] = make([]byte, m)
			fmt.Scan(&warehouse[j])
			for c := 0; c < m; c++ {
				if warehouse[j][c] == 'A' {
					A[0] = j
					A[1] = c
				} else if warehouse[j][c] == 'B' {
					B[0] = j
					B[1] = c
				}
			}
			// fmt.Printf("A: %d %d\nB: %d %d\n", A[0], A[1], B[0], B[1])
			// Построение маршрутов
		}

		if math.Sqrt(float64(A[0]*A[0]+A[1]*A[1])) < math.Sqrt(float64(B[0]*B[0]+B[1]*B[1])) {
			findPathStart(warehouse, A, "a")
			findPathEnd(warehouse, B, "b")
		} else {
			findPathStart(warehouse, B, "b")
			findPathEnd(warehouse, A, "a")
		}

		// Вывод результата
		for j := 0; j < n; j++ {
			fmt.Println(warehouse[j])
		}
	}
}
