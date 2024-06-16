package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	ai := make([][]int, t)
	p := make([]int, t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n, &p[i])

		ai[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&ai[i][j])
		}
	}
	for i := 0; i < t; i++ {

		var totalLostRevenue int64 = 0
		for j := 0; j < len(ai[i]); j++ {
			commission := float64(ai[i][j]) * float64(p[i])
			commission = math.Floor(commission)
			received := math.Floor(commission/100.0) * 100
			lostRevenue := commission - received
			totalLostRevenue += int64(lostRevenue)
		}

		fmt.Printf("%.2f\n", float64(totalLostRevenue)/100.0)
	}
}
