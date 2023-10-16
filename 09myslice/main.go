// If a slice is declared with no inputs, then by default, it is initialized as nil.
// Its length and capacity are zero.

package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// len and cap functions
	fmt.Printf("len=%d cap=%d \n", len(fruitList), cap(fruitList))

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 234)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	numbers1 := make([]int, 0, 5) //0 is len , 5 is cap
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
