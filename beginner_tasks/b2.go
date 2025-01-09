package beginner_tasks

import "fmt"

func B2() {
	// Take in two numbers
	var first int
	var second int

	fmt.Println("Enter your two numbers: ")
	fmt.Scanln(&first, &second)

	// Addition
	add := first + second
	fmt.Println("Addition: ", first, " + ", second, " = ", add)

	// Subtraction
	sub := first - second
	fmt.Println("Subtraction: ", first, " - ", second, " = ", sub)

	// Multiplication
	mul := first * second
	fmt.Println("Multiplication: ", first, " * ", second, " = ", mul)

	// Division
	div := first / second
	fmt.Println("Division: ", first, " / ", second, " = ", div)
}
