package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// MyHandler struct
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello new HTTP/2!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: &handler,
	}

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
	// If there is no SSL certificate the use server.ListenAndServer() and use
	// curl -I --http2 http://localhost:3000/ (without SSL)
	// curl -I --http2 --insecure https://localhost:3000/ (For SSL)
	// New to create our own SSL certificate
	// fmt.Println("Server Running....")
}
