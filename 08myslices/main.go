package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Understanding slices in Golang")

	var fruitList = []string{"Apple", "Orange", "Watermelon"}
	fmt.Println("FruitList is: ", fruitList)

	fruitList = append(fruitList, "Mango", "Grapes")

	fmt.Println("Fruit list after appending is: ", fruitList)

	fruitList = fruitList[1:5]
	fmt.Println("Fruit list after slicing is: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 92
	highScores[2] = 88
	highScores[3] = 76
	// highScores[4] = 77

	highScores = append(highScores, 70, 77, 65)

	fmt.Println(highScores)
	sort.Ints(highScores)

	fmt.Println("Sorted high scores: ", highScores)

	// How to remove a value from slices based on index
	var courses = []string{"java", "python", "golang", "javascript"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing: ", courses) // Prints: [java python javascript], Removes "golang" which was at position 2
}
