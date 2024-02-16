package router

import (
	"github.com/AliKhedmati/routate-backend/src/api/http/controllers"
	"github.com/AliKhedmati/routate-backend/src/database"
	"github.com/AliKhedmati/routate-backend/src/repositories"
	"github.com/AliKhedmati/routate-backend/src/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func init() {
	db = database.GetDatabase()
}

func NewRouter() *gin.Engine {

	// Initialize new router and attach logger and recovery middlewares.
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			/*
				Countries
			*/

			countries := v1.Group("/countries")
			{
				countries.GET("/list")
			}

			/*
				Cities
			*/

			cities := v1.Group("/cities")
			{
				cities.GET("")
			}

			/*
				Locations
			*/

			locations := v1.Group("/locations")
			{
				locations.GET("")
			}

			/*
				Users
			*/

			userRepo := repositories.NewUserRepository(db.Collection("users"))
			userService := services.NewUserService(userRepo)
			userController := controllers.NewUserController(userService)

			users := v1.Group("/users")
			{
				users.POST("", userController.CreateUser)
				users.GET(":id", userController.FindUserByID)
			}
		}
	}

	return router
}
