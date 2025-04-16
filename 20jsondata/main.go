package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Understanding more about JSON data in Golang")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"Go Programming", 299, "Udemy", "udemy123", []string{"high-level", "backend"}},
		{"ReactJS", 399, "Coursera", "react123", []string{"web-dev", "frontend"}},
		{"MERN Bootcamp", 499, "LearnCode", "learncode616", nil},
	}

	// Package this data as JSON data
	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
