package main

import (
	"fmt"
	"sync"
	"time"
)

func message1(ch1 chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("Message 1 value %d", i)
		time.Sleep(time.Millisecond * 2)
	}
	wg.Done()
}

func message2(ch2 chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		ch2 <- fmt.Sprintf("Message 2 value %d", i)
		time.Sleep(time.Millisecond * 2)
	}

	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)

	wg.Add(2)

	go message1(ch1, wg)
	go message2(ch2, wg)

	for i := 0; i < 10; i++ {
		select {
		case r := <-ch1:
			fmt.Println(r)
		case r := <-ch2:
			fmt.Println(r)
		}
	}

	wg.Wait()
	close(ch1)
	close(ch2)
}
