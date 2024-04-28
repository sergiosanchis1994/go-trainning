package main

import (
	"fmt"
	"math"
)

func main() {
	var side1, side2 float64
	fmt.Print("Enter side 1: ")
	fmt.Scanln(&side1)
	fmt.Print("Enter side 2: ")
	fmt.Scanln(&side2)

	hypotenuse := math.Sqrt(side1*side1 + side2*side2)
	area := (side1 * side2) / 2
	perimeter := side1 + side2 + hypotenuse

	fmt.Printf("Area: %f\n", area)
	fmt.Printf("Perimeter: %f\n", perimeter)
}
