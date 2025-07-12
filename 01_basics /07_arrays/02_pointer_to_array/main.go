package main

import "fmt"

func main() {
	originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// fmt.Println(originalArray)
	// fmt.Println(copiedArray)

	// // does't change the original array
	// copiedArray[0] = 100
	// fmt.Println(originalArray)
	// fmt.Println(copiedArray)

	// declared a pointer to a integer array of size 3
	var copiedArrayPtr *[3]int
	copiedArrayPtr = &originalArray

	fmt.Println("Original Array", originalArray)
	fmt.Println("Copied Array", *copiedArrayPtr)

	// can dynamically modify the original array via pointer
	copiedArrayPtr[0] = 202

	fmt.Println("Original Array", originalArray)
	fmt.Println("Copied Array", *copiedArrayPtr)

}
