package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

type PostTestSuite struct{}

// var db *sql.DB

func init() {
	// var err error
	// db, err = sql.Open("postgres", "user=gwp dbname=gwp password=123 sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	// mux.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.ID, Equals, 1)
}
