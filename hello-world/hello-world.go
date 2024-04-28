package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	fullName := "Sergio Sanchis \t (alias \"sergiosanchis1994\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "hello"
	fmt.Println(s[0])

	var r rune = '❤'
	fmt.Println(r)
}
