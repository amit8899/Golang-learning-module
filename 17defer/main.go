package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	// defer stores lines in stack
	// and line execution follows
	// LIFO

	defer fmt.Println("One")
	defer fmt.Println("Two")

	myDefer()
}

func myDefer() {
	fmt.Print("called")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
