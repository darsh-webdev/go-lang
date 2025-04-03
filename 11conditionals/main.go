package main

import "fmt"

func main() {
	fmt.Println("Understanding conditional statements in Golang")

	// If else statements

	loginCount := 23
	var result string

	// if loginCount > 10 {
	// 	result = "Welcome!"
	// }

	if loginCount > 10 {
		result = "Welcome!"
	} else {
		result = "Access denied"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than or equal to 10")
	}

}
