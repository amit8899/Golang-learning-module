package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	webList := []string{
		"https://google.com",
		"https://fb.com",
		"https://youtube.com",
		"https://go.dev",
	}

	for _, val := range webList {
		go getStatusCode(val)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(site string) {
	defer wg.Done()

	res, err := http.Get(site)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, site)
	}
}
