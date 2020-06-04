package main

import "fmt"

func main() {
	var hour, minute int

	fmt.Scan(&hour)
	fmt.Scan(&minute)
	calculateAlram(hour, minute)
}

func calculateAlram(hour, minute int) {
	minuteCarry := 0
	if minute < 45 {
		minuteCarry = 1
		minute = minute + 60 - 45
	} else {
		minute = minute - 45
	}

	if hour == 0 && minuteCarry == 1 {
		hour = 23
	} else if minuteCarry == 1 {
		hour = hour - 1
	}

	fmt.Printf("%d %d\n", hour, minute)
}
