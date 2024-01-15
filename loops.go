package main

import "fmt"

func main() {
	// Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Using a for loop to iterate over an array
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Using a for loop to iterate over a string
	text := "Hello"
	for _, char := range text {
		fmt.Println(char)
	}

	// Infinite loop with break statement
	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter == 3 {
			break
		}
	}
}