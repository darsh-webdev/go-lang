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
	DecodeJson()
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

func DecodeJson() {
	demoJsonData := []byte(`
	{
        "courseName": "Go Programming",
        "Price": 299,
        "website": "Udemy",
        "tags": ["high-level","backend"]
    }
	`)

	var myCourse course

	checkValid := json.Valid(demoJsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(demoJsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was invalid")
	}

	// Some cases where you just want to add data to key value pairs

	var myOnlineCourse map[string]interface{}
	json.Unmarshal(demoJsonData, &myOnlineCourse)
	fmt.Printf("%#v\n", myOnlineCourse)

	for k, v := range myOnlineCourse {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
