package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	var (
		helloWorldMessage string = quote.Hello()
		goMessage         string = quote.Go()
	)

	optMessage := quote.Opt()

	fmt.Println(helloWorldMessage)
	fmt.Println(goMessage)
	fmt.Println(optMessage)
}
