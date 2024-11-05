package repositories

import (
	"example/models"
	"strings"
)

type UserRepository interface {
	GetUserByID(id string) models.User
	GetUserByUsername(username string) models.User
	SetUser(user models.User) models.User
	DeleteUser(id string) models.User
}

type userRepository struct {
}

// User
var users []models.User

func GetUserByID(id string) models.User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return models.User{}
}

func GetUserByUsername(username string) models.User {
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user
		}
	}
	return models.User{}
}

func SetUser(user models.User) models.User {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
			return user
		}
	}
	return models.User{}
}

func DeleteUser(id string) models.User {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return u
		}
	}
	return models.User{}
}
