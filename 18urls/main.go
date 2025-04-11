package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://nocss.club/learn?isCssActive=false&currentMember=asdasda"

func main() {
	fmt.Println("Handling URL's in Golang")
	fmt.Println(myUrl)

	// Parsing
	result, _ := url.Parse(myUrl)

	fmt.Println("Url Scheme:", result.Scheme)
	fmt.Println("Url Host", result.Host)
	fmt.Println("Url Path", result.Path)
	fmt.Println("Url Params", result.RawQuery)
	fmt.Println("Url Port", result.Port())

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)

	fmt.Println(qparams["currentMember"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	// Creating a URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "nocss.club",
		Path:     "/learn",
		RawQuery: "lang=go",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
