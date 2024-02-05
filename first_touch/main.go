package main

import "fmt"

func main() {
	bmi()
}

func bmi() {
	height := 1.67
	weight := 65.0

	bmi := weight / (height * height)

	fmt.Println(bmi)
}

func hello() {
	fmt.Println("hello, go!")
}
