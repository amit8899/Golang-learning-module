package main

import "fmt"

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

}
