// If a slice is declared with no inputs, then by default, it is initialized as nil.
// Its length and capacity are zero.

package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

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

	// deleting element at given index
	highScores = deleteElement(highScores, 4)
	fmt.Println(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	numbers1 := make([]int, 0, 5) //0 is len , 5 is cap
	printSlice(numbers1)

	structSlice := []User{
		{Name: "Amit", Email: "mail", Status: true, Age: 56},
		{Name: "Dev", Email: "mail", Status: true, Age: 54},
	}
	fmt.Println("slice of struct:", structSlice)

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}

// delete single element
func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// delete multiple elements
func deleteElements(slice []int, indices []int) []int {
	// Sort indices in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, index := range indices {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

// The ... operator after slice[3:] is used to spread the
// elements in the slice. This means that instead of appending
// a slice to slice[:1], we append each individual element of
// slice[3:] to slice[:1]. This effectively removes the elements
// at index 1 and 2.
