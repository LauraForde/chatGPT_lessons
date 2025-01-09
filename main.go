package main

import (
	"fmt"

	"github.com/LauraForde/chatGPT_lessons/beginner_tasks"
)

var content = []string{
	"Hello World.",
	"Writing To File.",
}

func main() {

	var option int

	fmt.Println("Choose one of the following:" + "\n" + "1. Hello, World" + "\n" + "2. Basic Arithmetic Calculator" + "\n" + "3. FizzBuzz" +
		"\n" + "4. Temperature Converter" + "\n" + "5. Read and Write File")
	fmt.Scanf("%d", &option)

	switch {
	case option == 1:
		// Hello World
		beginner_tasks.B1()
	case option == 2:
		// Basic Arithmetic Calculator
		beginner_tasks.B2()
	case option == 3:
		// FizzBuzz
		beginner_tasks.B3()
	case option == 4:
		// Temperature Converter
		beginner_tasks.B4()
	case option == 5:
		// Read and Write Files
		beginner_tasks.B5()
	default:
		fmt.Println("Invalid option")
		return
	}

}
