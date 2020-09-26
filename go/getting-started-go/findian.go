package getting_started_go

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var value string

	fmt.Printf("Please enter a string value: ")
	untrimmedValue, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Printf("Input Error\n")
	} else {
		value = strings.TrimSuffix(untrimmedValue, "\n")
		lowerString := strings.ToLower(value)

		starts := strings.HasPrefix(lowerString, "i")
		ends := strings.HasSuffix(lowerString, "n")
		contains := strings.Contains(lowerString, "a")

		if starts && contains && ends {
			fmt.Printf("Found!\n")
		} else {
			fmt.Printf("Not Found!\n")
		}
	}
}
