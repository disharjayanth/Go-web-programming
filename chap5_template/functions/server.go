package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func formateDate(t time.Time) string {
	layout := "01-02-2006"
	fmt.Println(time.UnixDate)
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formateDate}
	t := template.New("t1.html").Funcs(funcMap)
	t, _ = t.ParseFiles("t1.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
