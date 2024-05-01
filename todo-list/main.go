package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	email string
}

func main() {
	p := Person{name: "Sergio", age: 29, email: "sergiosanchis1994@gmail.com"}
	p2 := Person{name: "Javier", age: 22, email: "javier@gmail.com"}
	p.age = 30
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p2)
}
