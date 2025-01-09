package beginner_tasks

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var content = []string{
	"Hello World.",
	"Writing To File.",
	"This is for question 5 of the beginners subset of questions I asked ChatGPT for to aid in my learning GO.",
}

func B5() {

	// Create file
	createF, err := os.Create("beginner5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer createF.Close()
	fmt.Println("File created: " + createF.Name())

	// Writing to file
	buffer := bufio.NewWriter(createF)

	for _, content := range content {
		_, err := buffer.WriteString(content + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}

	// Reading from file and printing to console
	fromFile, err := os.Open("beginner5.txt")
	if err != nil {
		fmt.Println(err)
	}

	readFile, err := io.ReadAll(fromFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Read from file: " + string(readFile))

}
