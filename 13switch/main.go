package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move start")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // further cases will also run
	case 4:
		fmt.Println("You can move 4 spot")
	default:
		fmt.Println("What was that!")
	}
}
