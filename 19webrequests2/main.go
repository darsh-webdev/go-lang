package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Handling web requests to local server in Golang")

	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post"

	// Fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Golang Masterclass",
			"price" : 0,
			"platform" : "learnGo.dev"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content for Post request is: ", string(content))

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform"

	// Form data
	data := url.Values{}
	data.Add("firstname", "Darshan")
	data.Add("lastname", "Rajput")
	data.Add("email", "darshan@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content for Post Form request is: ", string(content))
}
