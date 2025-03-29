package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Burger App!")
	fmt.Println("Please rate our burger between 1 and 5")

	var reader, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("You entered:", reader)

	var trimmed = strings.TrimSpace(reader) // This will remove the trailing newline character as it connot be directly converted to an integer

	var rating, err = strconv.ParseFloat(trimmed, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Rating:", rating+1)
	}
}
