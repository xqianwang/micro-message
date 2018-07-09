package routes

import (
	"github.com/micro-message/handlers"
)

//InitializeRoutes initialize routes for application
func initializeRoutes() {
	//Use SetUserStatus to indicate wheather the user request is authenticated or not
	router.Use(handlers.SetUserStatus())
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	//routes for users
	userRoutes := router.Group("/users")
	{
		//handle the GET request at /users/register
		//ensure that user is not logged in when request for user registration
		userRoutes.GET("/register", handlers.EnsureNotLoggedIn(), handlers.ShowResgistrationPage)
		//handle POST request at /users/register
		//ensure that user is not logged in when request for user registration
		userRoutes.POST("/register", handlers.EnsureNotLoggedIn(), handlers.Register)
		//handle GET request at /users/login
		//ensure that user is not logged in when request for user login
		userRoutes.GET("/login", handlers.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		//handle POST request at /users/login
		//ensure that user is not logged in when request for user login
		userRoutes.POST("/login", handlers.EnsureNotLoggedIn(), handlers.Login)
		//handle GET request at /users/logout
		//ensure that user is logged in when request for user logout
		userRoutes.GET("/logout", handlers.EnsureLoggedIn(), handlers.Logout)
	}

	//routes for messages
	messageRoutes := router.Group("/messages")
	{
		//handle GET request at /messages
		//this route will list messages
		messageRoutes.GET("", handlers.GetMessages)
		//handle GET request at /messages/view/:messageid
		//this route will show the message based on message id
		messageRoutes.GET("/view/:messageid", handlers.GetMessage)
		//handle DELETE request at /messages/:messageid
		//this route will delete message based on message id
		messageRoutes.DELETE("/:messageid", handlers.DeleteMessage)
		//handle POST request at /messages/create
		//this route will create message
		messageRoutes.POST("/create", handlers.CreateMessage)
		//handle GET request at /messages/create
		//this route will get create message page
		messageRoutes.GET("/create", handlers.ShowCreatePage)
	}
}
