package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Directory struct {
	Name    string      `json:"name"`
	Files   []string    `json:"files"`
	Subdirs []Directory `json:"subdirs"`
}

// func isInfected(d Directory) bool {
// 	for _, item := range d.Content {
// 		if item.IsInfected || strings.HasSuffix(item.Name, ".hack") {
// 			return true
// 		}
// 	}
// 	return false
// }

// func countInfectedFiles(dir Directory) int {
// 	count := 0
// 	for _, item := range dir.Content {
// 		if isInfected(item) {
// 			count++
// 		}
// 	}
// 	return count
// }

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		var root Directory

		str := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&str[j])
		}
		longStr := strings.Join(str, "")
		input := []byte(longStr)

		err := json.Unmarshal([]byte(input), &root)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			// return
		}

		// result := countInfectedFiles(root)
		fmt.Println(root)
	}
}
