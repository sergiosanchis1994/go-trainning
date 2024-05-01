package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)

	//target source
	elements := copy(slice2, slice1)

	fmt.Println("Elements:", elements)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
