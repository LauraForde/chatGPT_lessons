package main

import "fmt"

func main(){
	// Task 3 FizzBuzz

	// Using for/if loops
	/*for i := 1; i <= 100; i++{
		
		if i % 5 == 0 && i % 3 == 0{
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0  {
			fmt.Println("Buzz")
		} else{
			fmt.Println(i)
		}
	}*/

	// Using switch
	for i := 1; i <= 100; i++{
		switch{
		case(i % 5 == 0 && i % 3 == 0):
			fmt.Println("FizzBuzz")
		case(i % 3 == 0):
			fmt.Println("Fizz")
		case(i % 5 == 0):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}