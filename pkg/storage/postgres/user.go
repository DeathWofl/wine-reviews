package postgres

import (
	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) User(id uint) (*model.User, error) {
	var user *model.User
	s.DB.First(user, id)
	return user, nil
}

func (s *UserService) Users() ([]*model.User, error) {
	var users []*model.User

	s.DB.Find(&users)
	return users, nil
}

func (s *UserService) CreateUser(w *model.User) (*model.User, error) {
	s.DB.Create(w)
	return w, nil
}

func (s *UserService) DeleteUser(id uint) error {

	s.DB.Delete(&model.User{}, id)
	return nil
}

func (s *UserService) UpdateUser(id uint, w *model.User) error {
	var user *model.User
	s.DB.First(user, id)

	s.DB.Model(user).Updates(w)
	return nil
}
