package main

import (
	"html/template"
	"net/http"
)

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templ.html")
	t.Execute(w, r.FormValue("comment"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/form", form)
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
