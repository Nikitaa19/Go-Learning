package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)

	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	for i, value := range numbers {
		fmt.Printf("Index %d - Value %d\n", i, value)
	}

	slice1 := numbers[1:4]
	fmt.Println(slice1)

	slice1[0] = 10
	fmt.Println(numbers)

	slice2 := []string{"apple", "orange", "banana"}
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3}
	slice3 = append(slice3, 4, 5)
	fmt.Println(slice3)

	for i, value := range slice3 {
		fmt.Printf("Index %d - Value %d\n", i, value)
	}
}