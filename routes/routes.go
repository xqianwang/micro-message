package routes

import (
	"github.com/micro-message/handlers"
)

//InitializeRoutes initialize routes for application
func initializeRoutes() {
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/messages", handlers.GetMessages)

	//
	router.GET("/messages/view/:messageid", handlers.GetMessage)
	router.DELETE("/messages/:messageid", handlers.DeleteMessage)
	router.POST("/messages/create", handlers.CreateMessage)
	router.GET("/messages/create", handlers.ShowCreatePage)
}
