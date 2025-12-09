package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A int
	fmt.Scan(&A)

	s := strconv.Itoa(A)
	n := len(s)

	result := ""
	count := 0

	for i := n - 1; i >= 0; i-- {
		result = string(s[i]) + result
		count++

		if count == 3 && i != 0 {
			result = "," + result
			count = 0
		}
	}
	fmt.Println(result)
}
