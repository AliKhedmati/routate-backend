package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	// Initialize new router and attach logger and recovery middlewares.
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}
