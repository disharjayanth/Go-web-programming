package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func tearDown() {}

func TestHandlePost(t *testing.T) {
	jsonString := strings.NewReader(`{
		"content": "Hello there!" ,
		"author": "SauSheongChang"
	}`)
	request, _ := http.NewRequest("POST", "/post/", jsonString)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/3", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response Code is: %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)

	if post.ID != 3 {
		t.Errorf("Post id is: %v", post.ID)
	}

	if post.Content != "Hello there!" {
		t.Errorf("Post content is: %s", post.Content)
	}

	if post.Author != "SauSheongChang" {
		t.Errorf("Post author is: %s", post.Author)
	}
}

func TestHandleUpdate(t *testing.T) {
	jsonString := strings.NewReader(`{
		"content": "Updated Post Again" ,
		"author": "Sau Sheong Chang"
	}`)
	request, _ := http.NewRequest("PUT", "/post/1", jsonString)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}

func TestHandleDelete(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/post/4", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is: %v", writer.Code)
	}
}
