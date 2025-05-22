package main

import "fmt"

func main() {
	// Ex-1 Solution
	var i int = 20
	var f float64 = float64(i)

	fmt.Println("Value of i:", i)
	fmt.Println("Value of f:", f)

	// Ex-2 Solution
	const value = 10
	i1 := value
	var f1 float64 = value
	fmt.Println("Value of i1:", i1)
	fmt.Println("Value of f1:", f1)

	// Ex-3 Solution
	var b byte = 255
	var small1 int32 = 2147483647
	var big1 uint64 = 18446744073709551615

	b = b + 1
	small1 = small1 + 1
	big1 = big1 + 1

	fmt.Println("Value of b:", b)
	fmt.Println("Value of small1:", small1)
	fmt.Println("Value of big1:", big1)
}
