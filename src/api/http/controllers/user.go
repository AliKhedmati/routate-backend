package controllers

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/model"
	"github.com/AliKhedmati/routate-backend/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

// NewUserController initializes a new instance of UserController.
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

// CreateUser handles the creation of a new user.
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the user service
	if err := c.userService.Create(context.Background(), &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// FindUserByID finds a user by ID.
func (c *UserController) FindUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.FindByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
