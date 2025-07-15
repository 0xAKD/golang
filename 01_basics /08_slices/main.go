package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int                   //declaration of slice
	// var numbers1 = []int{1, 2, 3, 4, 5} //declaration of initialised slice

	// numbers2 := []int{6, 7, 8, 9, 10} //declaration of initialised slice without var keyword

	// slice := make([]int, 5)		//declared an empty slice using make.

	// making a slice with array elements
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4] // initialised slice with array values from index 1 to 4 i.e [2,3,4]

	slice1 = append(slice1, 7, 8, 9) // append more values into existing slice
	fmt.Println(slice1)

	// copying a slice into another one
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(slice1)

	slice1[2] = 40

	// check weather a slice is equal to another slice  or not
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("both the slices are equal!")
	} else {
		fmt.Println("the slices are not equal")
	}

	//	2D slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
