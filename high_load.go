package main

import (
	"fmt"
)

func isValidProcess(s string) bool {
	validStates := map[byte][]byte{
		'M': {'R', 'C', 'D'},
		'R': {'C'},
		'C': {'M'},
		'D': {'M'},
	}

	if len(s) <= 1 || s[0] != 'M' {
		return false
	}

	for i := 1; i < len(s)-1; i++ {
		if i > 0 && !contains(validStates[s[i-1]], s[i]) {
			return false
		}
	}
	if !contains(validStates[s[len(s)-1]], s[len(s)-2]) || s[len(s)-1] != 'D' {
		return false
	}

	return true
}

func contains(actions []byte, action byte) bool {
	for _, a := range actions {
		if a == action {
			return true
		}
	}
	return false
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)

		if isValidProcess(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
