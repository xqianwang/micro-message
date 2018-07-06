package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", ShowIndexPage)

	//Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Welcome to Micro Message!</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowCreatePage(t *testing.T) {
    r := getRouter(true)
    r.GET("/messages/create", ShowCreatePage)

    req, _ := http.NewRequest("GET", "/messages/create", nil)
    testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
        statusOK := w.Code == http.StatusOK

        p, err := ioutil.ReadAll(w.Body)
        pageOK := err == nil && strings.Index(string(p), "<title>Create Message</title>") > 0
        
        return statusOK && pageOK
    })
}

func TestGetMessages(t *testing.T) {
    r := getRouter(true)
    r.GET("/messages", GetMessages)

    req, _ := http.NewRequest("GET", "/messages", nil)
        testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
        statusOK := w.Code == http.StatusOK

        p, err := ioutil.ReadAll(w.Body)
        pageOK := err == nil && strings.Index(string(p), "<title>Message List</title>") > 0
        
        return statusOK && pageOK
    })
}
