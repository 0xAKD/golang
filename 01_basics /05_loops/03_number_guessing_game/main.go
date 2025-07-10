package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome Message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have choosen a numebr between 1 and 100")
	fmt.Println("Can you guess the number..??? ")

	var count int = 0
	var guess int
	for {
		fmt.Println("Enter your Guess: ")
		fmt.Scanln(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Invalid number!!")
			fmt.Println("Try guessing number between 1 to 100")
			continue
		}

		count++
		//check if the guess is correct
		if guess == target {
			fmt.Printf("Congrats! you guessed the number in %d guess", count)
			break
		} else if guess < target {
			fmt.Println("Too Low! Try guessing a Higher number ")
		} else {
			fmt.Println("Too High! Try guessing a Lower number ")
		}

	}

}
