package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Exploring file access in Golang")

	content := "This is random data to put in a file - LearningGo"

	file, err := os.Create("./demo.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./demo.txt")

}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
