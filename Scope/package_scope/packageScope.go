package main

import (
	"fmt"

	mathlib "example.com/mathLib" //Custom package calling
)

var (
	a = 50
	b = 20
)

/*
1. block -> {}
2. local scope -> defined within a block
3. global scope -> defined outside any block
4. go mod init example.com -> create a package
5. go mod tidy -> update the package
*/
func main() {
	fmt.Println("Showing package scope")
	/// When using another package in our package that's time we need used the function nmae in capital letter for access the function.
	mathlib.Add(a, b)
}
