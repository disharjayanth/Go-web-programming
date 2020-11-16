package main

import (
	"encoding/json"
	"fmt"
	"os"
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
				Content: "Have a great day sir",
				Author:  "Joe",
			},
			Comment{
				ID:      4,
				Content: "How are you sir?",
				Author:  "Rogan",
			},
		},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}

	fmt.Println("Exited.")
}
