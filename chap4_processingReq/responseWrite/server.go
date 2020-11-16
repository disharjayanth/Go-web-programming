package main

import (
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str :=
		`<html>
		<head><title>Go Web Programming</title></head>
		<body><h1 style="color: blue;">Hello World</h1></body>
		</html>`

	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/process", writeExample)

	server.ListenAndServe()
}
