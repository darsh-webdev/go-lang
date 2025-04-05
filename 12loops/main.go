package main

import "fmt"

func main() {
	fmt.Println("Understanding loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days:", days)

	// For loop
	for d := 0; d < len(days); d++ {
		fmt.Printf("Day at position %v: %v\n", d, days[d])
	}

	// For using range
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Day at position %v: %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 8 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Rogue value is:", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jump to position 9")
}
