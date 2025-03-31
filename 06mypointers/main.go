package main

import "fmt"

func main() {
	fmt.Println("Understanding pointers in Golang")

	// var myPointer *int
	// fmt.Println("Value of pointer is ", myPointer)

	myNumber := 23
	var ptr = &myNumber // ptr now "knows the address" of myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual variable is", *ptr) // "Looking inside the address" - prints 23

	*ptr = *ptr + 2                       // "Putting new data in the address" - changes myNumber to 25
	fmt.Println("New value is", myNumber) // Prints 25
}
