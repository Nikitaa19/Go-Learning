package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alice",
		"age":     "25",
		"country": "USA",
	}

	fmt.Println(person)

	fmt.Println(person["name"])
	fmt.Println(person["age"])

	person["age"] = "26"
	fmt.Println(person["age"])

	person["gender"] = "Female"
	fmt.Println(person)

	delete(person, "country")
	fmt.Println(person)

	country, exists := person["country"]
	if exists {
		fmt.Println("Country:", country)
	} else {
		fmt.Println("Country key does not exist")
	}
}