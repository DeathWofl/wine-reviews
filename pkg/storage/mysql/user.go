package mysql

import (
	"github.com/deathwofl/wine-reviews/pkg"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) User(id uint) (*pkg.User, error) {

	var user *pkg.User
	s.DB.First(user, id)
	return user, nil
}

func (s *UserService) Users() []*pkg.User {
	var users []*pkg.User

	s.DB.Find(users)
	return users
}

func (s *UserService) CreateUser(w *pkg.User) error {

	s.DB.Create(w)
	return nil
}

func (s *UserService) DeleteUser(id uint) error {

	s.DB.Delete(&pkg.User{}, id)
	return nil
}

func (s *UserService) UpdateUser(id uint, w *pkg.User) error {
	var user *pkg.User
	s.DB.First(user, id)

	s.DB.Model(user).Updates(w)
	return nil
}
