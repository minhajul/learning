package main

import "fmt"

func main() {
	maps := make(map[string]int)
	fmt.Println(maps)

	maps["a"] = 1
	maps["b"] = 2
	maps["c"] = 1
	fmt.Println(maps)

	delete(maps, "a")

	fmt.Println(maps)

	fmt.Println("Length: ", len(maps))
}
