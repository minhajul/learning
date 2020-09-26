package getting_started_go

import "fmt"

func main() {
	var floatNumber float64

	fmt.Printf("Please enter a float number: ")
	num, err := fmt.Scan(&floatNumber)

	trunc := int(floatNumber)

	if num == 1 && err == nil {
		fmt.Printf("Your integer is: %d\n", trunc)
	} else {
		fmt.Printf("Input Error\n")
	}
}
