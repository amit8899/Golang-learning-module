package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
	Refer bool
}

func main() {

	User1 := &User{"Amit", "mail.com", 24, true}
	jm, _ := json.Marshal(User1)
	fmt.Println(string(jm))

	var user2 User
	by := `{"Name":"Amit","Email":"mail.com","Age":28,"Refer":true}`
	json.Unmarshal([]byte(by), &user2)

	fmt.Printf("%+v", user2)

}
