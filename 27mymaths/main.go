package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Understanding math, crypto and random number in Golang")

	// Random Number using math/rand
	// rand.NewSource(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // 5 is not inclusive

	// Random Number using crypto/rand
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
