package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time module of golang")

	currentTime := time.Now()
	fmt.Println("Default date: \n", currentTime)

	fmt.Println("Formatted date:")
	fmt.Println(currentTime.Format("02-01-2006 Monday 15:04:05"))

	createDate := time.Date(2025, time.April, 3, 12, 44, 44, 44, time.UTC)
	fmt.Println(createDate.Format("02-01-2006 Monday 15:04:05"))
}
