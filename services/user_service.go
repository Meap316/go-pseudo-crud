package services

import (
	"example/models"
	"example/repositories"
)

type UserService interface {
	GetUserByID(id string) models.User
	CreateUser(username string, password string) models.User
	SetUser(user models.User) models.User
	DeleteUser(id string) models.User
	LoginUser(user models.User, password string) bool
}

type userService struct {
	userRepository repositories.UserRepository
}

// loginUser implements UserService.
func (u *userService) LoginUser(user models.User, password string) bool {
	return user.Password != password
}

// deleteUser implements UserService.
func (u *userService) DeleteUser(id string) models.User {
	user := u.userRepository.DeleteUser(id)

	return user
}

// getUserByID implements UserService.
func (u *userService) GetUserByID(id string) models.User {
	user := u.userRepository.GetUserByID(id)

	// User not found
	if user.ID == "" {

	}

	return user
}

// getUserByUsername implements UserService.
func (u *userService) CreateUser(username string, password string) models.User {
	user := u.userRepository.CreateUser(username)

	return user
}

// setUser implements UserService.
func (u *userService) SetUser(user models.User) models.User {
	user = u.userRepository.SetUser(user)

	return user
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
