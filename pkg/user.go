package pkg

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type UserService interface {
	User(id string) (*User, error)
	Users() []*User
	CreateUser(w *User) error
	DeleteUser(id string) error
	UpdateUser(id string, w *User) error
}
