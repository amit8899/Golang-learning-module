package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	// languages := make(map[string]string)  //1

	// var languages = map[string]string{} //2

	var languages = map[string]string{"Go": "golang", "RS": "rust"} //3

	// adding items
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// delete items
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	users := map[int]User{
		1: {"Amit", 15},
		2: {"Pankaj", 25},
	}

	us2 := User{"Shrey", 24}
	users[4] = us2

	fmt.Println(users)

}
