package getting_started_go

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var stringValue []int = make([]int, 3)
	var inputValue string
	fmt.Println("Please enter an integer ('X' to Exit):")
	for true {
		fmt.Scanln(&inputValue)
		if inputValue == "X" {
			break
		}
		value, err := strconv.Atoi(inputValue)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		stringValue = append(stringValue, value)
		sort.Ints(stringValue[:])
		fmt.Println(stringValue)
		fmt.Println("Please enter an integer ('X' to Exit):")
	}
}
