package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = [...]int{10, 20, 30, 40, 50}
	a[0] = 100
	a[1] = 200
	fmt.Println("Matrix", a)

	for i := 0; i < len(a); i++ {
		fmt.Println("Matrix position "+strconv.Itoa(i)+": ", a[i])
	}

	for index, value := range a {
		fmt.Printf("Index: %d, Value %d\n", index, value)
	}

	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)

}
