package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `form:"username"`
	Password  string `form:"password"`
	IsDeleted bool
}
