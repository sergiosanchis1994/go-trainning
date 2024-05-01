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

	fmt.Println(colors["red"])
	colors["black"] = "#000000"
	fmt.Println(colors)
	value, ok := colors["red"]
	fmt.Println(value, ok)
	if value, ok = colors["white"]; ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println(value, ok)
	delete(colors, "red")
	fmt.Println(colors)
}
