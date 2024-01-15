package main

import "fmt"

func main() {
	// Integers
	var age int = 25
	var temperature int32 = 30

	fmt.Println("Age:", age)
	fmt.Println("Temperature:", temperature)

	// Floating-Point Numbers
	var price float64 = 49.99
	var discount float32 = 0.1

	fmt.Println("Price:", price)
	fmt.Println("Discount:", discount)

	// Strings
	var name string = "Alice"
	var greeting = "Hello"

	fmt.Println(greeting, name)

	// Booleans
	var isRaining bool = true
	var isSunny = false

	fmt.Println("Is it raining?", isRaining)
	fmt.Println("Is it sunny?", isSunny)

	// Constants
	const pi = 3.14
	const gravity float32 = 9.81

	fmt.Println("Pi:", pi)
	fmt.Println("Gravity:", gravity)
}