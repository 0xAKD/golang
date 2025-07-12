package main

import "fmt"

func main() {
	fruit := "Mango"

	switch fruit {
	case "Apple":
		fmt.Println("It's Apple")
	case "Banana":
		fmt.Println("It's Banana")
	case "Mango":
		fmt.Println("It's Mango")
		fallthrough // execute the next case also even if the condition is false
	case "Orange":
		fmt.Println("It's Orange")
	default: // if no case matches,the default statement executes
		fmt.Println("None of the above")
	}

	// multiple conditions
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday":
		fmt.Println("It's a Weekday")
	case "Sunday":
		fmt.Println("It's a Weekend")
	default:
		fmt.Println("Invalid day")
	}

	checktype(0200)
	checktype(3.14)
	checktype("World")
	checktype(true)
}

// function to check the datatype of a value.
func checktype(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}

}
