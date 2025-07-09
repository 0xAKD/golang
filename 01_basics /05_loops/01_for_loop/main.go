package main

import "fmt"

func main() {

	// print numbers from 1 to 10
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	// iteration over a collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// for range loop
	for i := range 10 { // prints numbers from 1 to 10
		fmt.Println(i + 1)
	}

	// loop control statements
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			continue // skips the iteration
		}
		fmt.Println("Odd Number: ", i)
		if i == 5 {
			break // breaks the loop
		}
	}
}
