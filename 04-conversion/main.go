package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza app"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for Pizza:")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 2, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
