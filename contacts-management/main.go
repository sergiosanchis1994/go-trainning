package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	_, err = file.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
