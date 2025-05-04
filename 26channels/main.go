package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Understanding channels in Golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
