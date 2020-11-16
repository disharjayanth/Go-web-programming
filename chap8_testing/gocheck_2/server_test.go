package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
	post   *FakePost
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test:", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite.")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finished Post Test Suite:")
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.ID, Equals, 1)
}

func (s *PostTestSuite) TestHandlePut(c *C) {
	jsonReader := strings.NewReader(`{
		"content": "Updated Post" ,
		"author": "SauSheongChang"
	}`)

	request, _ := http.NewRequest("PUT", "/post/1", jsonReader)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.ID, Equals, 1)
	c.Check(s.post.Content, Equals, "Updated Post")
	c.Check(s.post.Author, Equals, "SauSheongChang")
}
