package main

import "fmt"

func main() {
	var d CustomError
	d.data = "Can't assign value"
	fmt.Println(d.Error())
}

type CustomError struct {
	data string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occured due to %v", e.data)
}
