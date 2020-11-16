package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templ.html")
	t.Execute(w, "Helllooo..???")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
