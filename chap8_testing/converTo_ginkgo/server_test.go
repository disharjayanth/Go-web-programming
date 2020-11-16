package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("get post", func() {

		mux := http.NewServeMux()
		mux.HandleFunc("/post/", handleRequest(&FakePost{}))

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/post/1", nil)
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is: %v", writer.Code)
		}
		var post Post
		json.Unmarshal(writer.Body.Bytes(), &post)
		if post.ID != 1 {
			GinkgoT().Errorf("Post ID is: %v", post.ID)
		}
	})
})
