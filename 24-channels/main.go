package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-myCh

		fmt.Println(isChanOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
