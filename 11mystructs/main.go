package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent

	amit := User{"Amit", "mail", true, 16}
	fmt.Println(amit)
	fmt.Printf("amit details: %+v", amit)
	fmt.Printf("Name is %v and email is %v.", amit.Name, amit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
