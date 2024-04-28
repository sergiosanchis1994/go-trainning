package main

import (
	"fmt"

	"rsc.io/quote"
)

const Pi float32 = 3.14

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(Pi)
	fmt.Println(Wednesday)
}
