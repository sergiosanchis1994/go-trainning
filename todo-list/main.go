package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#00FF00",
	}

	for key, value := range colors {
		fmt.Printf("%s: %s\n", key, value)
	}
}
