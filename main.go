package main

import (
	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `uri:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
}

// User
var users []User

func main() {

	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
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
	})

	r.POST("/user", func(c *gin.Context) {
		// Check if username exist
		var user = getUserByUsername(c.PostForm("username"))

		if user.Username != "" {
			c.JSON(400, gin.H{
				"status":  "failed",
				"message": "Username already exist",
			})
			return
		}

		// Create user
		user = User{
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
	})

	r.PUT("/user/:id", func(c *gin.Context) {
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
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
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
	})

	r.POST("/login", func(c *gin.Context) {
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
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getUserByID(id string) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

func getUserByUsername(username string) User {
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user
		}
	}
	return User{}
}

func setUser(user User) User {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
			return user
		}
	}
	return User{}
}

func deleteUser(user User) User {
	for i, u := range users {
		if u.ID == user.ID {
			users = append(users[:i], users[i+1:]...)
			return user
		}
	}
	return User{}
}
