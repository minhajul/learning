package main

import (
	"fmt"
)

func main() {
	negativePositive(1)

	a, b := values(1, 3)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(compareNumber(3, 5))
}

func compareNumber(numberOne, numberTwo int) bool {
	return numberOne > numberTwo
}

func values(numberOne, numberTwo int) (int, int){
	return numberOne, numberTwo
}

func negativePositive(num int) {
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}