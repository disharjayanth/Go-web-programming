package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}

	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	// statement := "insert into posts (content, author) values ($1, $2) returning id"
	// stmt, err := Db.Prepare(statement)
	// if err != nil {
	// 	return
	// }
	// defer stmt.Close()
	// err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	// return
	// OR ->
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{
		Content: "Hello World!",
		Author:  "Sau Sheong Chang",
	}

	fmt.Println("After initialisation post of type Post:", post)
	post.Create()
	fmt.Println("After creating the row and storing into Post table:", post)

	readPost, _ := GetPost(post.Id)
	fmt.Println("Get the post having the id:", post.Id, "->", readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println("Getting all posts of limit 10", posts)

	readPost.Delete()

	posts, _ = Posts(10)
	fmt.Println("Getting all posts after deleting one of the post:", posts)

	fmt.Println("Exited.")
}
