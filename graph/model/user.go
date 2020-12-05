package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Reviews   []*Review `json:"reviews"`
}

type UserService interface {
	User(id uint) (*User, error)
	UserbyEmail(email string) (*User, error)
	UserbyUsername(username string) (*User, error)
	Users() ([]*User, error)
	CreateUser(w *User) (*User, error)
	DeleteUser(id string) error
	UpdateUser(id string, w *User) error
}
