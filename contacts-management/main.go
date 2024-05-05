package main

import (
	"fmt"
)

func divide(dividend, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("You cannot divide by 0")
	}
}

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)
}
