package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	jsonString := strings.NewReader(`{
		"content": "New Post" ,
		"author": "SauSheongChang"
	}`)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/post/", jsonString)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.ID != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	jsonString := strings.NewReader(`{
		"content": "Updated Post" ,
		"author": "SauSheongChang"
	}`)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/post/1", jsonString)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}

func TestHandleDelete(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/post/4", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}
