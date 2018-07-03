package handlers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro-message/store"
)

//ShowResgistrationPage shows registration page
func ShowResgistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register",
	}, "register.html")
}

//Register registers new user
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	user, err := store.RegisterUser(username, password, email)
	if err != nil || user == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	} else {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{
			"title": "Successful registration & Login"}, "register-successful.html")
	}
}

//ShowLoginPage performs showing login page
func ShowLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

//Login performs user login
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if store.ValidUser(username, password, email) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

//Logout performs user logout
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
