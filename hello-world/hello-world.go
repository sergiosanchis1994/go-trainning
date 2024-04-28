package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32)) //150

	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i) //200

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s) //4242

}
