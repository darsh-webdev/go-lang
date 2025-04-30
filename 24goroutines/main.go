package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually these are pointers

func main() {
	fmt.Println("Understanding Goroutines in Golang")
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://google.in",
		"https://instagram.com",
		"https://github.com",
		"https://go.dev",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // Always at the end of the main function
}

func greeter(s string) {
	for range 6 {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
