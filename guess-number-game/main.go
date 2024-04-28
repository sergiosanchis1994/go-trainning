package main

import "fmt"

func main() {
	fmt.Println(hello("Sergio"))
	sum, prod := calc(4, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Prod:", prod)
}

func hello(name string) string {
	return "Hello " + name
}

func calc(a, b int) (sum, prod int) {
	sum = a + b
	prod = a * b
	return
}
