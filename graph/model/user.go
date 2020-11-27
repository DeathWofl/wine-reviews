package model

type User struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Reviews  []*Review `json:"reviews"`
}

type UserService interface {
	User(id string) (*User, error)
	Users() []*User
	CreateUser(w *User) error
	DeleteUser(id string) error
	UpdateUser(id string, w *User) error
}
