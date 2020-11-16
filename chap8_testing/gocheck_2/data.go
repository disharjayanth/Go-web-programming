package main

import (
	"database/sql"
)

type Post struct {
	Db      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Text interface {
	retrieve(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

func (post *Post) retrieve(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	err = post.Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.ID)
	return
}

func (post *Post) update() (err error) {
	_, err = post.Db.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = $1", post.ID)
	return
}
