package main

import "fmt"

// sum function takes any number of integers and returns their sum
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", result)
}