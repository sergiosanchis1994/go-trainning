package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	var name string
	var age int

	fmt.Print("What's your name?")
	fmt.Scanln(&name)

	fmt.Print("How old are you?")
	fmt.Scanln(&age)

	fmt.Printf("Hi, my name is %s and I've %d years old\n", name, age)

	greeting := fmt.Sprintf("Hi, my name is %s and I've %d years old", name, age)
	fmt.Println(greeting)

}
