package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}
