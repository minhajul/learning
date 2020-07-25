package main

import "fmt"

type Post struct {
	title       string
	author      string
	isPublished bool
}

func main() {
	post := &Post{title: "This is dummy post", author: "Minhajul"}
	makePublish(post)
	fmt.Println(post)
}

func makePublish(p *Post) {
	p.isPublished = true
}
