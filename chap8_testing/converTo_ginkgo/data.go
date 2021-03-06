package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Pointer if post will implement Text interface
type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Get a single post
func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	statment := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := post.Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update a given post
func (post *Post) update() (err error) {
	_, err = post.Db.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

// Delete a given post
func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = $1", post.ID)
	return
}
