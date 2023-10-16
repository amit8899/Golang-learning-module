package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent

	amit := User{"Amit", "mail", true, 16}
	fmt.Println(amit)
	fmt.Printf("amit details: %+v\n", amit)
	fmt.Printf("Name is %v and email is %v.\n", amit.Name, amit.Email)
	amit.GetStatus()
	amit.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "new-email"
	fmt.Println("Email of this user: ", u.Email)
}
