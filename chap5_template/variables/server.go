package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html")
	someData := map[int]string{
		1: "John",
		2: "Doe",
		3: "Smith",
		4: "Brad",
		5: "Ken",
		6: "Ray",
	}
	t.Execute(w, someData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
