package main

import "fmt"

var (
	a = 100
	b = 200
)

func printNumber(num int) {
	fmt.Println("Number is:", num)
}

func add(a int, b int) {
	res := a + b
	printNumber(res)
}

func main() {
	fmt.Println("Showing boring example of scope")
	add(a, b)
}
