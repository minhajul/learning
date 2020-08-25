package main

import "fmt"

func main() {
	var rows int
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
