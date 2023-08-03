package main

import "fmt"

func main() {
	weekDay("Sunday")
}



func weekDay(day string) {
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	case "Thursday":
		fmt.Println("It's Thursday.")
	case "Friday":
		fmt.Println("It's Friday.")
	default:
		fmt.Println("It's weekend")
	}
}
