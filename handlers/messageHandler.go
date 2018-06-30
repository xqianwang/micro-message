package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro-message/store"
)

//ShowIndexPage shows index page and message list inside it
func ShowIndexPage(c *gin.Context) {
	messages, _ := store.GetAllMessages()
	//fresh index page according to the values
	render(c, gin.H{
		"title":   "Welcome to Qlik Message center",
		"payload": messages}, "index.html")
}

//ShowCreatePage shows create message page
func ShowCreatePage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create Message"}, "create-message.html")
}

//GetMessages shows list of messages and
func GetMessages(c *gin.Context) {
	if messages, err := store.GetAllMessages(); err == nil {
		render(c, gin.H{
			"title":   "Message List",
			"payload": messages}, "messages.html")
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//GetMessage get message based on message id
func GetMessage(c *gin.Context) {
	//convert parameter value from string to integer
	if messageId, err := strconv.Atoi(c.Param("messageid")); err == nil {
		if message, err := store.GetMessageByID(messageId); err == nil {
			render(c, gin.H{
				"id":      message.ID,
				"payload": message}, "message.html")
		} else {
			//if message not found
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid message id"))
	}
}

//CreateMessage ceates a new message
func CreateMessage(c *gin.Context) {
	// Obtain the POSTed content values
	content := c.PostForm("content")
	if id, err := store.CreateMessage(content); err == nil {
		render(c, gin.H{
			"title": "Submission Successful",
			"id":    id}, "submission.html")
	} else {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to create message"))
	}
}

//DeleteMessage deletes an old message
func DeleteMessage(c *gin.Context) {
	if messageId, err := strconv.Atoi(c.Param("messageid")); err == nil {
		if err := store.DeleteMessage(messageId); err == nil {
			c.Writer.WriteHeader(http.StatusOK)
		} else {
			c.AbortWithError(http.StatusNoContent, fmt.Errorf("Not find corresponding message based in id %d", messageId))
		}
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
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
