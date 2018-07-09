package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestRegistrationPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/users/register", ShowResgistrationPage)

	req, err := http.NewRequest("GET", "/users/register", nil)
	if err != nil {
		t.Error(err)
	}

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		OK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") >= 0

		return OK && pageOK
	})
}

func TestNewRegister(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)
	r.POST("/users/register", Register)

	newUser := getRegistrationUser()
	req, err := http.NewRequest("POST", "/users/register", strings.NewReader(newUser))
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(newUser)))

	r.ServeHTTP(w, req)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		OK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Successful registration & Login</title>") >= 0

		return OK && pageOK
	})
}

func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/users/login", ShowLoginPage)

	req, _ := http.NewRequest("GET", "/users/login", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

func TestLoginUnauthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.POST("/users/login", Login)
	loginPayload := getLoginUser()
	req, _ := http.NewRequest("POST", "/users/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err == nil || strings.Index(string(p), "<title>Successful Login</title>") > 0 {
		t.Fail()
	}
}

func TestLoginAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.POST("/users/login", Login)
	loginPayload := getLoginUser()
	req, _ := http.NewRequest("POST", "/users/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful Login</title>") < 0 {
		t.Fail()
	}
}
