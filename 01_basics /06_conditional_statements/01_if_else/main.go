package main

import "fmt"

func main() {

	// simple example of if-else statement
	num := 15
	if num > 10 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// nested if statements
	if num > 10 {
		if num > 12 {
			fmt.Println("Both conditions are true")
		} else {
			fmt.Println("Only the first condition is true")
		}
	} else {
		fmt.Println("None of the conditions are true")
	}

	// && AND
	if num > 10 && num < 20 {
		fmt.Println("Both conditions are true")
	} else {
		fmt.Println("False")
	}

	// || OR
	if num > 10 || num < 20 {
		fmt.Println("Either of the conditions are true")
	} else {
		fmt.Println("False")
	}

	// Else-if
	if num < 10 {
		fmt.Println("First condition is true")
	} else if num > 12 {
		fmt.Println("Second Condition is true")
	} else {
		fmt.Println("None of the above conditions are true")
	}

}
