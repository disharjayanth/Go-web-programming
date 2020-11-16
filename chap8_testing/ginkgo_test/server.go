package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=123 sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Listening PORT: @3000")

	http.HandleFunc("/post/", HandleRequest(&Post{Db: db}))

	server.ListenAndServe()
}

// HandleRequest func
func HandleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err = post.retrieve(id)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}

	err = post.update()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}
	err = post.create()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.retrieve(id)
	if err != nil {
		return
	}
	jsonData, err := json.MarshalIndent(post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = post.retrieve(id)
	if err != nil {
		return
	}

	err = post.delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}
