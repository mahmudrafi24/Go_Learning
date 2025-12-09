package main

import "fmt"

var (
	a = 50
	b = 20
)

/*
1. block -> {}
2. local scope -> defined within a block
3. global scope -> defined outside any block
4. variable shadowing -> inner block variable with same name as outer block variable
*/
func main() {
	x := 18

	if x >= 18 {
		p := 10
		fmt.Println("I am matured boy")
		fmt.Println("I've ", p, " years of experience")
	}
	// fmt.Println("I've ", p, " years of experience") // p is not defined outside the block

	fmt.Println("I've ", a, " years of experience") // p is not defined outside the block

}
