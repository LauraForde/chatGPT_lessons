package beginner_tasks

import "fmt"

func B1() {
	// Task 1
	//fmt.Println("Hello, World!")

	// Task 2
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")
}
