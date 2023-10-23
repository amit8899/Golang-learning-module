/*
Error is generated by you as the Go coder, when you expect something to go wrong and when it did,
you will return an error. Panic is when the error is unexpected and the program cannot continue on.
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var d CustomError
	d.data = "Can't assign value"
	fmt.Println(d.Error())

	result, err := Sqrt(-1)

	fmt.Println("square root of given num is:", result)

	if err != nil {
		fmt.Println(err)
	}
}

type CustomError struct {
	data string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occured due to %v", e.data)
}

// function returning error
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil

}
