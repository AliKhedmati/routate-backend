package controllers

import (
	"github.com/AliKhedmati/routate-backend/src/api/http/responses"
	"github.com/AliKhedmati/routate-backend/src/models"
	"github.com/AliKhedmati/routate-backend/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController represents a controller for user-related operations.
type UserController struct {
	userService *services.UserService
}

// NewUserController returns an instance of UserController.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Create handles the creation of a new user.
func (c *UserController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the user services
	if err := c.userService.Create(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": &user})
}

// FindByID finds a user by ID.
func (c *UserController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.userService.FindByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Message: "OK!",
		Data: map[string]interface{}{
			"user": &user,
		},
	})
}
