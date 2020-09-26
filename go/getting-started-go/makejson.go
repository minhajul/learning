package getting_started_go

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var person map[string]string
	person = make(map[string]string)
	fmt.Println("Please input your name:")

	br := bufio.NewReader(os.Stdin)
	byteName, _, _ := br.ReadLine()
	name := string(byteName)

	person["name"] = name

	fmt.Println("Then input your address:")
	brAgain := bufio.NewReader(os.Stdin)
	byteAddress, _, _ := brAgain.ReadLine()
	address := string(byteAddress)

	person["address"] = address

	b, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Encoding Field")
	} else {
		fmt.Println("Encoded Data: ")
		fmt.Println(b)
		fmt.Println("Decoded Data: ")
		fmt.Println(string(b))
	}
}
