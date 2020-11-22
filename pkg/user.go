package pkg

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserService interface {
	User(id string) (*User, error)
	Users() []*User
	CreateUser(w *User) error
	DeleteUser(id string) error
	UpdateUser(id string, w *User) error
}
