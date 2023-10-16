package main

import "fmt"

const LoginToken string = "bhjhhoihoi" // first letter capital shows public variable

func main() {
	var username string = "Amit"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallInt int8 = 127
	fmt.Println(smallInt)
	fmt.Printf("variable is of type: %T \n", smallInt)

	var smallFloat float32 = 255.687576576567
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)
}
