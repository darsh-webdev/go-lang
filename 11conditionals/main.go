package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Understanding conditional statements in Golang")

	//<---------------------- If else statements ----------------------->
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

	//<---------------------- Switch statement ----------------------->
	source := rand.NewSource(time.Now().UnixNano()) // seed with a specific value
	r := rand.New(source)
	diceNumber := r.Intn(6) + 1

	fmt.Println("Value of dice is:", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("Invalid dice roll")
	}
}
