package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 21:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Night")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Unknown")
	}
}
