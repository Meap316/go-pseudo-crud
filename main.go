package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {

	// Open Database Connection
	dsn := "host=68.183.179.230 user=admin password=0qP2JDV3EVkEMEd dbname=latihan_gorm_irvan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Database Migration

	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		// Get user
		var user = getUserByID(c.Param("id"), db)

		// User not found
		if user.ID == 0 {
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
		var user = getUserByUsername(c.PostForm("username"), db)

		if user.Username != "" {
			c.JSON(400, gin.H{
				"status":  "failed",
				"message": "Username already exist",
			})
			return
		}

		// Create user

		user = createUser(c.PostForm("username"), c.PostForm("password"), db)

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
		var user = getUserByID(c.Param("id"), db)

		// User not found
		if user.ID == 0 {
			c.JSON(404, gin.H{
				"status":  "failed",
				"message": "User not found",
			})
			return
		}

		// Update user
		user.Username = c.PostForm("username")
		user.Password = c.PostForm("password")

		setUser(user, db)

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
		var user = getUserByID(c.Param("id"), db)

		// User not found
		if user.ID == 0 {
			c.JSON(404, gin.H{
				"status":  "failed",
				"message": "User not found",
			})
			return
		}

		deleteUser(user, db)

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

		var user = getUserByUsername(username, db)

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

func createUser(username string, password string, db *gorm.DB) User {

	user := User{
		Username: username,
		Password: password,
	}

	db.Create(&user)

	return user
}

func getUserByID(id string, db *gorm.DB) User {

	user := User{}

	db.First(&user, "id = ?", id)

	return user
}

func getUserByUsername(username string, db *gorm.DB) User {

	user := User{}

	db.First(&user, "username = ?", username)

	return user
}

func setUser(user User, db *gorm.DB) User {

	db.Save(&user)

	return user
}

func deleteUser(user User, db *gorm.DB) User {

	db.Delete(&user)

	return User{}
}
