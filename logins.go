package main

import (
	"fmt"
)

// Helper function to check if two strings are similar
func areSimilar(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	res := true
	diffCount := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
			if diffCount > 1 {
				res = false
				break
			}
			if i < len(s1)-1 {
				if s1[i+1] == s2[i+1] {
					res = false
					break
				} else if s1[i] != s2[i+1] {
					res = false
					break
				}
				i++
			} else {
				res = false
				break
			}
		}
	}

	return res
}
func getSum(str string) int64 {
	var sum int64 = 0
	for c := 0; c < len(str); c++ {
		sum += int64(str[c])
	}
	return sum * int64(len(str))
}

func main() {
	var n, m int
	fmt.Scan(&n)

	cl := make(map[int64][]string)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		sum := getSum(str)
		if cl[sum] == nil {
			cl[sum] = make([]string, 0)
		}
		cl[sum] = append(cl[sum], str)
	}

	fmt.Scan(&m)

	newLogins := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&newLogins[i])
	}

	for _, newLogin := range newLogins {
		sum := getSum(newLogin)
		res := false
		for i := 0; i < len(cl[sum]); i++ {
			if areSimilar(newLogin, cl[sum][i]) {
				fmt.Println(1)
				res = true
				break
			}
		}
		if res == false {
			fmt.Println(0)
		}

	}
}
