package main

import "fmt"

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	for {
		fmt.Println("hello world")
	}

}
