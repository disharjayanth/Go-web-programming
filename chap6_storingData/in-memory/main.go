package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostsById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostsById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostsById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{
		Id:      1,
		Content: "Hello World!",
		Author:  "Sau Sheong Chang",
	}

	post2 := Post{
		Id:      2,
		Content: "Bonjour Monde!",
		Author:  "Pierre",
	}

	post3 := Post{
		Id:      3,
		Content: "Hola Mundo!",
		Author:  "Pedro",
	}

	post4 := Post{
		Id:      4,
		Content: "Greetings Earthlings",
		Author:  "Sau Sheong Chang",
	}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println("Individual access:")
	fmt.Println(PostsById[1])
	fmt.Println(PostsById[2])
	fmt.Println(PostsById[3])
	fmt.Println(PostsById[4])

	fmt.Println("Ranging over PostsById map:")
	for id, post := range PostsById {
		fmt.Println(id, "->", *post)
	}

	fmt.Println("List of posts from author: Sau Sheong Chang ->")
	for _, post := range PostsByAuthor["Sau Sheong Chang"] {
		fmt.Println(post)
	}

	fmt.Println("List of posts from author: Pierre ->")
	for _, post := range PostsByAuthor["Pierre"] {
		fmt.Println(post)
	}
}
