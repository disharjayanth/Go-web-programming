package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Post struct
type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello world!</h1></body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "Service not available!")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com/")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	post := Post{
		User:    "sausheong chang",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
