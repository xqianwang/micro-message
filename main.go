package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro-message/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//Get router
	router := routes.InitRouter()
	//run the router applciation
	router.Run()
}
