package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	randomNumber := rand.Intn(100)
	var num int
	var tries int
	const maxTries = 10

	for tries < maxTries {
		tries++
		fmt.Printf("Introduce a number (tries: %d): ", maxTries-tries+1)
		fmt.Scanln(&num)

		if num == randomNumber {
			fmt.Printf("Congrats! %d was the secret number!\n", num)
			newGame()
			return
		} else if num < randomNumber {
			fmt.Println("Your number is lower than the secret number")
		} else if num > randomNumber {
			fmt.Println("Your number is bigger than the secret number")
		}
	}
	fmt.Printf("You lost :( Secret number was: %d \n", randomNumber)
	newGame()
}

func newGame() {
	var choice string
	fmt.Printf("Do you want to play again? (y/n): ")
	fmt.Scanln(&choice)
	switch choice {
	case "y":
		play()
	case "n":
		fmt.Println("Bye :)")
	default:
		fmt.Println("Invalid option. Please, introduce only \"y\" or \"n\"")
		newGame()
	}
}
