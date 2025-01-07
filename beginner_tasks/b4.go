package main

import "fmt"

func main(){
	// Task 4 convert celsius to fahrenheit

	var cel float64
	var fah float64

	fmt.Println("Enter temp in Celsius: ")
	fmt.Scanln(&cel)

	fah = (cel * 1.8) + 32
	fmt.Println("Temp in Fahrenheit: ", fah)

}