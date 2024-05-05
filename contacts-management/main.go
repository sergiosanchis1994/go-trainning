package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main(): ")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hi! I'm a log message!")
}
