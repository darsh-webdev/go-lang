package main

import "fmt"

func main() {
	fmt.Println("Understanding maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages are: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Languages after deleting: ", languages)

	// Loops in golang (Looping over a map)
	for key, value := range languages {
		fmt.Printf("For key %v, value is: %v\n", key, value)
	}
}
