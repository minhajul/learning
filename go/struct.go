package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Minhajul", 1000})

	personObj := person{name: "Sean", age: 50}

	fmt.Println(personObj.name)
}
