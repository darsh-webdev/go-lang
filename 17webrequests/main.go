package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://nocss.club/"

func main() {
	fmt.Println("Handling web requests in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is: %T\n", response)

	defer response.Body.Close() // Caller's responsibity to close this connnection, as ReadResponse nor Response.Write never closes a connection.

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("Content is: ", content)
}
