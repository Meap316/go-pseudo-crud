package controllers

import (
	"example/models"
	"example/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) getUserByID(c *gin.Context) {
	// Get user
	userId := c.Param("id")

	user := uc.userService.GetUserByID(userId)

	// return user
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "User found",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	})
}

func (uc *UserController) createUserByUsername(c *gin.Context) {
	// Check if username exist
	username := c.PostForm("username")

	user := getUserByUsername()

	if user.Username != "" {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "Username already exist",
		})
		return
	}

	// Create user
	user = models.User{
		ID:       strconv.Itoa(len(users) + 1),
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	users = append(users, user)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "User created",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	})
}

func (uc *UserController) setUser(c *gin.Context) {
	// Get user
	var user = getUserByID(c.Param("id"))

	// User not found
	if user.ID == "" {
		c.JSON(404, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	// Update user
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")

	setUser(user)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "User updated",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	})
}

func (uc *UserController) deleteUser(c *gin.Context) {
	// Get user
	var user = getUserByID(c.Param("id"))

	// User not found
	if user.ID == "" {
		c.JSON(404, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	deleteUser(user)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "User deleted",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	})
}

func (uc *UserController) login(c *gin.Context) {
	// Login

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check if user exist

	var user = getUserByUsername(username)

	if user.Username == "" {
		c.JSON(404, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	if user.Password != password {
		c.JSON(404, gin.H{
			"status":  "failed",
			"message": "Wrong password",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Login success",
	})
}
