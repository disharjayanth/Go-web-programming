package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func main() {
	log.SetOutput(os.Stdout)
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/headers", headers)
	log.Println("Listening at PORT 3000")
	server.ListenAndServe()
}
