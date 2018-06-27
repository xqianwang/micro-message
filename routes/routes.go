package routes

import (
	"github.com/micro-message/handlers"
)

//InitializeRoutes initialize routes for application
func initializeRoutes() {
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/message/view/:messageid", getMessage)
}
