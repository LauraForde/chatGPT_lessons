package main

import "fmt"

func main(){
	// Task 1
	//fmt.Println("Hello, World!")

	// Task 2
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")
}