package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent

	amit := User{"Amit", "mail", true, 16}
	fmt.Println(amit)
	fmt.Printf("amit details: %+v", amit)
	fmt.Printf("Name is %v and email is %v.\n", amit.Name, amit.Email)

	var amit2 = User{Name: "Amit", Status: true, Age: 30}
	fmt.Println(amit2)

	var amit3 = new(User) //instance of User
	amit3.Name = "Pankaj"
	amit3.Status = false
	fmt.Println(amit3)

	amit4 := new(User) //pointer to instance
	amit4.Name = "Amit"
	fmt.Println(amit4)

	// Struct Instantiation Using Pointer Address Operator
	amit5 := &User{"devesh", "mail@", true, 65}
	fmt.Println(amit5)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
