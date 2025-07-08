package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition = ", result)

	result = a - b
	fmt.Println("Addition = ", result)

	result = a * b
	fmt.Println("Addition = ", result)

	result = a / b
	fmt.Println("Addition = ", result)

	result = a % b
	fmt.Println("Addition = ", result)

	//for getting the decimal values, one of the operands must be float type
	const pi = 22 / 7.0
	fmt.Println("Value of PI = ", pi)

	//Overflow with signed integers
	var maxInt int64 = 9223372036854775807 //max value that int64 can hold
	fmt.Println("max value that int64 can hold: ", maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//Overflow with signed integers
	var uMaxInt uint64 = 18446744073709551615 //max value that int64 can hold
	fmt.Println("max value that uint64 can hold: ", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	//Underflow with signed integers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
