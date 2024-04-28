package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)
}
