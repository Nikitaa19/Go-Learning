package main

import "fmt"

// Numeric Constants
const (
	pi       = 3.14
	gravity  = 9.81
	avogadro = 6.022e23
)

// String Constants
const (
	appName  = "MyApp"
	version  = "1.0"
	welcome  = "Welcome to"
	greeting = welcome + " " + appName + " " + "version " + version
)

// Boolean Constants
const (
	isProduction = true
	isDebug      = false
)

// iota and Enum-like Constants
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// Numeric Constants
	fmt.Println("Pi:", pi)
	fmt.Println("Gravity:", gravity)
	fmt.Println("Avogadro's Number:", avogadro)

	// String Constants
	fmt.Println(greeting)

	// Boolean Constants
	fmt.Println("Is Production Environment?", isProduction)
	fmt.Println("Is Debug Mode?", isDebug)

	// Enum-like Constants
	fmt.Println("Days of the Week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
}