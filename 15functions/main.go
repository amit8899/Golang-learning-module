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

	var (
		number1 int = 2
		number2 int = 6
	)

	// getting multiple returns from same function
	addition, subtraction, multiplication, division := addSubMulDiv(number1, number2)

	fmt.Printf("The addition of %d and %d is %d.\n", number1, number2, addition)
	fmt.Printf("The subtraction of %d and %d is %d.\n", number1, number2, subtraction)
	fmt.Printf("The multiplication of %d and %d is %d.\n", number1, number2, multiplication)
	fmt.Printf("The division of %d and %d is %d.\n", number1, number2, division)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(val1, val2 int) int {
	return val1 + val2
}

// multiple arguments function
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi! Pro result function"
}

func addSubMulDiv(number1, number2 int) (int, int, int, int) {
	var addition, subtraction, multiplication, division int

	// adding two numbers
	addition = number1 + number2

	// subtracting two numbers
	subtraction = number1 - number2

	// multiplying two numbers
	multiplication = number1 * number2

	// dividing two numbers
	division = number1 / number2
	return addition, subtraction, multiplication, division
}
