package main

import "fmt"

func main() {
	year := 2020

	if isLeapYear(year) {
		fmt.Println("LeapYear")
	} else {
		fmt.Println("NotLeapYear")
	}
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
