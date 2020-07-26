package main

import (
	"fmt"
)

type Post struct {
	title       string
	author      string
	isPublished bool
}

func main() {
	post := &Post{title: "This is dummy post", author: "Minhajul"}
	makePublish(post)
	fmt.Println(post)

	posts_array := make([]Post, 2)
	posts_array = append(posts_array, Post{
		title: "This is dummy post", author: "Minhajul", isPublished: true,
	})

	posts_array = append(posts_array, Post{
		title: "This is dummy post", author: "Minhajul", isPublished: false,
	})

	fmt.Println(getPublishedPost(posts_array))
}

func makePublish(p *Post) {
	p.isPublished = true
}

func getPublishedPost(posts []Post) []string {
	posts := make([]string, 2)
	for index, post := range posts {
		if post.isPublished {
			posts[index] = post.title
		}
	}

	return posts
}
