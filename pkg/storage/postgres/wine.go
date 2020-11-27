package postgres

import (
	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

type WineService struct {
	DB *gorm.DB
}

func (s *WineService) Wine(id uint) (*model.Wine, error) {

	var wine *model.Wine
	s.DB.First(wine, id)
	return wine, nil
}

func (s *WineService) Wines() []*model.Wine {
	var wines []*model.Wine

	s.DB.Find(wines)
	return wines
}

func (s *WineService) CreateWine(w *model.Wine) error {

	s.DB.Create(w)
	return nil
}

func (s *WineService) DeleteWine(id uint) error {

	s.DB.Delete(&model.Wine{}, id)
	return nil
}

func (s *WineService) UpdateWine(id uint, w *model.Wine) error {
	var wine *model.Wine
	s.DB.First(wine, id)

	s.DB.Model(wine).Updates(w)
	return nil
}
