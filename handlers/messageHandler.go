package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro-message/models"
)

func ShowIndexPage(c *gin.Context) {
	messages := models.GetAllMessages()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": messages,
		},
	)
}

func GetMessageByID(id int) {
	var message models.Message
	for 
}
