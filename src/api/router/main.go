package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	// Initialize new router and attach logger and recovery middlewares.
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("")
			}

			countries := v1.Group("/countries")
			{
				countries.GET("")
				countries.POST("")
			}
		}
	}

	return router
}
