package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - website.com"

	// creating file
	file, err := os.Create("./mylocfile.txt")

	if err != nil {
		panic(err)
	}

	//writing file
	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close() //may be we write further after it

	// readFile("./mylocfile.txt")
	readFile(file.Name())

	appendFile(file.Name())

	readFile(file.Name())
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text inside file: ", string(databyte))
}

// append in a file
func appendFile(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		//   fmt.Println("Could not open example.txt")
		checkNilErr(err)
		return
	}

	defer file.Close()

	_, err2 := file.WriteString("\nAppending some text to example.txt")

	if err2 != nil {
		// fmt.Println("Could not write text to example.txt")
		checkNilErr(err)

	} else {
		fmt.Println("Operation successful! Text has been appended to ", filename)
	}
}

func checkNilErr(err error) {
	if err != nil {
		// error formator function
		err := fmt.Errorf("error is %v", err.Error())
		fmt.Println(err)

		panic(err)
	}
}
