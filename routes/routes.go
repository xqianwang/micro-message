package routes

import (
	"github.com/micro-message/handlers"
)

//InitializeRoutes initialize routes for application
func initializeRoutes() {
	//check status
	router.Use(handlers.SetUserStatus())
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	//routes for users
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/register", handlers.EnsureNotLoggedIn(), handlers.ShowResgistrationPage)
		userRoutes.POST("/register", handlers.EnsureNotLoggedIn(), handlers.Register)
		userRoutes.GET("/login", handlers.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", handlers.EnsureNotLoggedIn(), handlers.Login)
		userRoutes.GET("/logout", handlers.EnsureLoggedIn(), handlers.Logout)
	}

	messageRoutes := router.Group("/messages")
	{
        messageRoutes.GET("", handlers.EnsureLoggedIn(), handlers.GetMessages)
		messageRoutes.GET("/view/:messageid", handlers.EnsureLoggedIn(), handlers.GetMessage)
        messageRoutes.DELETE("/:messageid", handlers.EnsureLoggedIn(), handlers.DeleteMessage)
		messageRoutes.POST("/create", handlers.EnsureLoggedIn(), handlers.CreateMessage)
        messageRoutes.GET("/create", handlers.EnsureLoggedIn(), handlers.ShowCreatePage)
        messageRoutes.GET("/delete", handlers.EnsureLoggedIn(), handlers.ShowCreatePage)
	}
}
