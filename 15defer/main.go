package main

import "fmt"

func main() {
	fmt.Println("Understanding defer in Golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

// Will be printed as Hello, Two, One, World. (LIFO for deferred)

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
