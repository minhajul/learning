package main

import "fmt"

type person struct {
	name string
	age int
	isAdult bool
}

func main()  {
	someone := &person{name: "Minhaj", age: 50}
	checkAdult(someone)
	fmt.Println(someone)
}

func checkAdult(p *person)  {
	p.isAdult = p.age >= 18
}