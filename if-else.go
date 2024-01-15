package main

import "fmt"

func main() {
	//Basic if statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	//if-else statement
	y := 3
	if y%2 == 0 {
		fmt.Println("y is even")
	} else {
		fmt.Println("y is odd")
	}

	//Multiple if-else if-else statements
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	//if statement with short variable declaration
	if z := 15; z > 10 {
		fmt.Println("z is greater than 10")
	}

	//Nested if statements
	a := 5
	b := 10
	if a > 0 {
		if b > 5 {
			fmt.Println("Both a and b are positive")
		} else {
			fmt.Println("a is positive, but b is not")
		}
	} else {
		fmt.Println("a is not positive")
	}
}