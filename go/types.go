package main

import "fmt"

type User struct {
	id int
	name, location string
}

type Player struct {
	User
	game_id int
}

func main()  {
	fmt.Println("Learning composition and interface")
	player := Player{}
	player.id = 1
	player.name = "Minhajul"
	player.location = "Dhaka"
	fmt.Printf("%+v", player)

	arrays := [2]int{1, 2}
	fmt.Println(arrays)

	stringArray := [...]string{"Hello", "Kabiraj", "aksndkad"}
	fmt.Printf("%s is ", stringArray)
	fmt.Println(len(stringArray))

	var a[2][2]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}

	fmt.Printf("%q", a)
}
