package main

import "fmt"

func main() {
	// while loops example
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	// Do-while loops example
	j := 1
	for {
		fmt.Println("Example 6:", j)
		j++
		if j > 5 {
			break
		}
	}
}