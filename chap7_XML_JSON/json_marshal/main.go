package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		ID:      1,
		Content: "Hello World!",
		Author: Author{
			ID:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				ID:      3,
				Content: "Nice one",
				Author:  "Adam",
			},
			Comment{
				ID:      4,
				Content: "Good one",
				Author:  "John",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println(string(output))
	fmt.Println("Exited.")
}
