package main

import "fmt"

func main() {
	message := "Hello, Go!"

	length := len(message)
	fmt.Println("Length of the message:", length)

	firstChar := message[0]
	fmt.Println("First character:", string(firstChar))

	name := "Alice"
	greeting := "Hello, " + name + "!"
	fmt.Println(greeting)

	multiline := `
        This is a
        multi-line
        string.
    `
	fmt.Println(multiline)

	age := 25
	formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedString)
}