package main

import "fmt"

func main() {
	fmt.Println("Understanding structs in Golang")
	// No inheritance in golang, No super or parent or child

	darshan := User{"Darshan", "test@go.dev", true, 22}

	fmt.Println(darshan)
	fmt.Printf("darshan details are: %+v\n", darshan)

	fmt.Printf("Name is %v and age is %v.", darshan.Name, darshan.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
