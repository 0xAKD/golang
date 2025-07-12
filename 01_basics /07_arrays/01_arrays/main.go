package main

import "fmt"

func main() {
	var numArray [5]int   // declared an uninitialised array of type int
	fmt.Println(numArray) // all the members will be zero as the array is uninitialised.

	for i := 0; i < 5; i++ {
		numArray[i] = (i + 1) * 10 // initialising each element of the array
	}

	fmt.Println(numArray)

	// length of arrays
	fmt.Println("Length of array is: ", len(numArray))

	// initialised array
	fruits := [4]string{"Apple", "Banana", "Mango", "Pineapple"}

	// the underscore (_) is a blank identifier, its used when we dont want to use or store that value
	for _, v := range fruits {
		fmt.Printf("Value: %s", v)
	}

}
