package main 
/// Must be needed a package at every go file

import "fmt"

/// Main go file -> package will be main package
func main() {
	// Variable 
	/* 
	int --> int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
	float --> float32, float64
	complex --> complex64, complex128
	bool --> bool
	string --> string
	rune --> rune
	byte --> byte
	*/
	// var variable_name data_type = value
	fmt.Println("Hello world!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints a message to the console.\n\n")

	fmt.Println("Printing variable : int")
	var x int = 20
	var y int = 10
	var z int = x + y
	fmt.Println("The Sum of the two numbers is: %d\n", z)

	fmt.Println("Declaring variable : float64")
	var a float64 = 20.36
	var b float64 = 10.25
	var c float64 = a + b
	fmt.Println("The Sum of the two numbers is: %f ", c)

	fmt.Println("Declaring variable : string")
	rafi := "Rafi"
	fmt.Printf("My name is %s\n", rafi)

	fmt.Println("Declaring variable : complex64")
	var m complex64 = 5 + 5i
	var n complex64 = 10 + 10i
	var o complex64 = m + n
	fmt.Println("The Sum of the two numbers is: %", o)

	fmt.Println("Declaring variable : bool")
	var p bool = true
	var q bool = false
	fmt.Println("The value of p is: ", p)
	fmt.Println("The value of q is: ", q)

	var r string = "Hello"
	var s string = "World"						
	var t string = r + " " + s						
	fmt.Println("The sum of the two strings is: ", t)


	fmt.Println("Hello world!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints a message to the console.")
}