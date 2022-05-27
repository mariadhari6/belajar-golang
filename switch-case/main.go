package main

import "fmt"

func main() {
	number := 4
	var day string
	switch number {
	case 1:
		day = "Monday"
	case 2:
		day = "Tuesday"
	case 3:
		day = "Wednesday"
	case 4:
		day = "Thursday"
	case 5:
		day = "Friday"
	case 6:
		day = "Saturday"
	case 7:
		day = "Sunday"
	default:
		day = "DEFAULT"
	}
	fmt.Println(day)
}
