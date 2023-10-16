package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")

	greeter()

	result := adder(3, 6)
	fmt.Println("Result is: ", result)

	proRes, msg := proAdder(5, 58, 6, 10)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("pro message: ", msg)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(val1, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi! Pro result function"
}
