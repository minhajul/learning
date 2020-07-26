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

	postsArray := make([]Post, 2)
	postsArray = append(postsArray, Post{
		title: "This is dummy post", author: "Minhajul", isPublished: true,
	})

	postsArray = append(postsArray, Post{
		title: "This is dummy post", author: "Minhajul", isPublished: false,
	})

	fmt.Println(getPublishedPost(postsArray))
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
