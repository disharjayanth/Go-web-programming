package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	XMLName xml.Name `xml:"author"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "Hello World",
		Author: Author{
			ID:   "4",
			Name: "It was good.",
		},
	}

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Cannot create xmlfile:", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(post)
	if err != nil {
		fmt.Println("Cannot encode struct to xml to file:", err)
		return
	}
}
