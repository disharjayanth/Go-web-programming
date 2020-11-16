package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.PostFormValue("hello"), r.PostFormValue("post"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
