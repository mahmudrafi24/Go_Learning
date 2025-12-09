package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Println(i)
		}
	}
}
