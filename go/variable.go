package main

import (
	"fmt"
	"math"
)

// Struct
type Artist struct {
	name, genre string
	songs int
}

func newRelease(artist *Artist) int{
	artist.songs++
   return artist.songs
}

func main() {
	//name, age, location := "Minhajul", 50, "Dhaka"
	//fmt.Println("Name is %s age is %d from %s", name, age, location)
	//fmt.Printf("%s is also known as %s", name, location)
	fmt.Println("My favorite number is", math.Sqrt(2))
	fmt.Println(add(1, '2'))
	//
	region, continent := location("Santa Monica")
	fmt.Printf("%s and %s", region, continent)
	fmt.Println("/")

	// call newRelease
	me := &Artist{name:"Minhaj", genre:"Ajaira", songs:100}
	fmt.Printf("Artist name is %s and he sang total %d songs", me.name, newRelease(me))
	fmt.Printf("Artist name is %s and he sang total %d songs", me.name, me.songs)
}

func add(x, y int) int {
	return x + y
}

func location(city string) (string, string){
	var region, continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}