package main

import "fmt"

func add(num1 int, num2 int){
	sum := num1 + num2
	fmt.Println(sum)
}
func subtract(num1 int, num2 int){
	subtract := num1 - num2
	fmt.Println(subtract)
}
func multiply(num1 int, num2 int){
	multiply := num1 * num2
	fmt.Println(multiply)
}
func divide(num1 int, num2 int){
	divide := num1 / num2
	fmt.Println(divide)
}

func returnFunciton(num1 int, num2 int) (int, int, int, int){
	sum := num1 + num2
	subtract := num1 - num2
	multiply := num1 * num2
	divide := num1 / num2
	return sum, subtract, multiply, divide
}

func printSomething(){
	fmt.Println("Education must be free")
}

func sayHello(name string){
	fmt.Println("Hello", name)

}
func detailsFunction(){
	fmt.Println("👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌👌")
	fmt.Println("Welcome to the Application")
	var name string
	fmt.Println("Enter your name => ")
	fmt.Scanln(&name)

	var num1, num2 int
	fmt.Println("Enter first number => ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number => ")
	fmt.Scanln(&num2)

	sum := num1 + num2

	fmt.Println("Hello", name)
	fmt.Println("Sum of two numbers is =>", sum)
	fmt.Println("Thank you for using this application")
}
func main() {
	a := 10
	b := 20

	add(a, b)
	subtract(a, b)
	multiply(a, b)
	divide(a, b)	

	sum, subtract, multiply, divide := returnFunciton(a, b)
	fmt.Println("Return Function - Sum:", sum, "Subtract:", subtract, "Multiply:", multiply, "Divide:", divide)

	printSomething()
	sayHello("Rafi")

	detailsFunction()
}