package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro-message/models"
)

//ShowIndexPage shows index page and message list inside it
func ShowIndexPage(c *gin.Context) {
	messages := models.GetAllMessages()
	//fresh index page according to the values
	render(c, gin.H{
		"title":   "Welcome to Qlik Message center",
		"paylaod": messages,
	}, "index.html")
}

//Get message based on message id
func GetMessage(c *gin.Context) {
	//convert parameter value from string to integer
	if messageId, err := strconv.Atoi(c.Param("messageid")); err == nil {
		if message, err := models.GetMessageByID(messageId); err == nil {
			c.HTML(
				http.StatusOK,
				"message.html",
				gin.H{
					"id":      message.ID,
					"payload": message,
				},
			)
		} else {
			//if message not found
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid message id"))
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
