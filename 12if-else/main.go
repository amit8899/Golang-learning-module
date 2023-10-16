package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	LoginCount := 23
	var result string

	if LoginCount < 10 {
		result = "Regular user"
	} else if LoginCount > 10 {
		result = "something else"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// if err != nil {

	// }
}
