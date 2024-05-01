package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	email string
}

func (p *Person) hello() {
	fmt.Println("Hello, my name is", p.name)
}

func main() {
	var x int = 10
	fmt.Println(x)
	edit(&x)
	fmt.Println(x)

	p := Person{"Sergio", 30, "sergiosanchis1994@gmail.com"}
	p.hello()
}

func edit(x *int) {
	*x = 20
}
