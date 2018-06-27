package routes

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var once sync.Once

//InitRouter initialize gin router
func InitRouter() *gin.Engine {
	once.Do(func() {
		//set route as gin default one
		router = gin.Default()
		//load templates into gin context
		router.LoadHTMLGlob("templates/*")
		// Initialize the routes
		initializeRoutes()

	})
	return router
}
