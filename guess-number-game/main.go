package main

import (
	"fmt"
	"time"
)

func main() {
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Morning")
	} else if t.Hour() < 21 {
		fmt.Println("Afternoon")
	} else {
		fmt.Println("Night")
	}
}
