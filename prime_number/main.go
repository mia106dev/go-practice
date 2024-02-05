package main

import (
	"fmt"
	"math"
)

func main() {
	printIsPrime(1)
	printIsPrime(2)
	printIsPrime(9)
	printIsPrime(10)
	printIsPrime(13)
}

func printIsPrime(num int) {
	fmt.Print(num);

	if isPrime(num) {
		fmt.Println(" is Prime.")
	} else {
		fmt.Println(" is Not Prime.")
	}
}

// NOTE: 負の数を想定しない（定義外）
func isPrime(num int) bool {
	if num == 2 {
		return true
	} else if num < 2 || num%2 == 0 {
		return false
	}

	sqNum := int(math.Sqrt(float64(num)))

	for i := 3; i <= sqNum; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
