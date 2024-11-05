package models

type User struct {
	ID       string `uri:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
}
