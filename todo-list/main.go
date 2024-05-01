package main

import (
	"fmt"
)

func main() {
	week := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	days := week[0:5]

	fmt.Println(week)
	fmt.Println(days)

	days = append(days, "Friday")
	days = append(days, "Saturday")
	days = append(days, "Extra day")
	fmt.Println(days)

	days = append(days[:2], days[3:]...)
	fmt.Println(days)
	fmt.Println("Length", len(days))
	fmt.Println("Cap", cap(days))
}
