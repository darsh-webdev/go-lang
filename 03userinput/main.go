package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for this course:")

	// Comma ok || err
	input, _ := read.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is: %T \n", input)
}
