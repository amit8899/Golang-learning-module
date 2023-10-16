package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"

	fmt.Println("fruit List: ", fruitList)

	// Initializing an Array with an Array Literal

	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)

	// Initializing an Array with ellipses...

	x1 := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x1).Kind())
	fmt.Println(len(x1))

	// Initializing Values for Specific Elements

	x2 := [5]int{1: 10, 3: 30}
	fmt.Println(x2)

}
