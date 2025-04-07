package main

import "fmt"

func main() {
	fmt.Println("Understanding methods in Golang")

	darshan := User{"Darshan", "test@go.dev", true, 22}

	fmt.Println(darshan)
	fmt.Printf("darshan details are: %+v\n", darshan)
	fmt.Printf("Name is %v and age is %v.\n", darshan.Name, darshan.Age)

	darshan.GetStatus()
	darshan.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", darshan.Name, darshan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "darshan@go.dev"
	fmt.Println("Eamil of this user is: ", u.Email)
}
