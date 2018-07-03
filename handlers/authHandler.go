package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//SetUserStatus sets flag is_logged_in  based on checking token cookie in context
func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

//EnsureLoggedIn checks if is_logged_in is set or not. If it is not set, it will abort with HTTP auth error
func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

//EnsureNotLoggedIn checks if is_logged_in is set or not. If it is set, it will abort with HTTP auth error
func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
