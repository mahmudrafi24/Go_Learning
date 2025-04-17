package main

import "fmt"

func main() {
	age := 10

	money := 100000

	if age < 18 {
		fmt.Println("You are not eligible to Married")
	} else if age > 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are eligible to vote")
	}

	if age > 18 && money > 100000 {
		fmt.Println("You are eligible to married and have enough money")
	} else if age > 18 && money < 100000 {
		fmt.Println("You are eligible to married but you need more money")
	} else if age < 18 && money > 100000 {
		fmt.Println("You are not having aged to married but you have enough money")
	} else if age < 18 && money < 100000 {
		fmt.Println("You have enough age to married and you also don't have enough money")
	}

	if money > 100000 {
		fmt.Println("You have enough money to getting married")
	} else if money == 100000 {
		fmt.Println("You have money but you need to more money to getting married")
	} else {
		fmt.Println("You have money but getting married is risker decision to you")
	}

	//  Greate than => ">"
	//  Less than => "<"
	//  Greate than or equal to => ">="
	//  Less than or equal to => "<="
	//  Equal to => "=="
	//  Not equal to => "!="
	//  Logical operators => "&&" and "||"
	//  Logical AND => "&&"
	//  Logical OR => "||"
	//  Logical NOT => "!"
	//  Logical AND => "&&" => true && true = true
	//  Logical OR => "||" => true || false = true
	//  Logical NOT => "!" => !true = false
	//  Logical AND => "&&" => true && false = false
	//  Logical OR => "||" => false || false = false
	//  Logical NOT => "!" => !false = true
}
