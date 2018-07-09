package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

//TestEnsursLoggedIn tests EnsureLoggedIn function
func TestEnsureLoggedIn(t *testing.T) {
	testCases := map[string]struct {
		loggedIn bool
		code     int
	}{
		"Test function EnsureLoggedIn without authentication": {
			false,
			http.StatusUnauthorized,
		},
		"Test function EnsureLoggedIn with authentication": {
			true,
			http.StatusOK,
		},
	}

	for caseName, testCase := range testCases {
		t.Run(caseName, func(t *testing.T) {
			r := getRouter(false)
			r.GET("/", setLogIn(testCase.loggedIn), EnsureLoggedIn(), func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			testAuthHandler(t, r, testCase.code)
		})
	}
}

//TestEnsursNotLoggedIn tests  EnsureNotLoggedIn function
func TestEnsureNotLoggedIn(t *testing.T) {
	testCases := map[string]struct {
		loggedIn bool
		code     int
	}{
		"Test function EnsursNotLoggedIn without authentication": {
			false,
			http.StatusOK,
		},
		"Test function EnsursNotLoggedIn with authentication": {
			true,
			http.StatusUnauthorized,
		},
	}

	for caseName, testCase := range testCases {
		t.Run(caseName, func(t *testing.T) {
			r := getRouter(false)
			r.GET("/", setLogIn(testCase.loggedIn), EnsureNotLoggedIn(), func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			testAuthHandler(t, r, testCase.code)
		})
	}
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)
}

func setLogIn(loggedIn bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", loggedIn)
	}
}
