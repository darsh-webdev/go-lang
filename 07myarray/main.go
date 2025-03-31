package main

import "fmt"

func main() {
	fmt.Println("Understanding arrays in Golang")

	var fruitList [4]string // Declares an array with size of 4

	fruitList[0] = "Apple"
	fruitList[1] = "Pineapple"
	fruitList[3] = "Watermelon"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList)) // This will always print 4, even though only 3 elements are there in the list, this happens because during declaration the length of the array was mentioned as 4.

	var birdList = [3]string{"Pigeon", "Crow", "Parrot"} // Decalres and initializes and array with the provided values
	fmt.Println("Bird list is: ", birdList)
}
