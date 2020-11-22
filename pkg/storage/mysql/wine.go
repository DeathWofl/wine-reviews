package mysql

import (
	"github.com/deathwofl/wine-reviews/pkg"
	"gorm.io/gorm"
)

type WineService struct {
	DB *gorm.DB
}

func (s *WineService) Wine(id uint) (*pkg.Wine, error) {

	var wine *pkg.Wine
	s.DB.First(wine, id)
	return wine, nil
}

func (s *WineService) Wines() []*pkg.Wine {
	var wines []*pkg.Wine

	s.DB.Find(wines)
	return wines
}

func (s *WineService) CreateWine(w *pkg.Wine) error {

	s.DB.Create(w)
	return nil
}

func (s *WineService) DeleteWine(id uint) error {

	s.DB.Delete(&pkg.Wine{}, id)
	return nil
}

func (s *WineService) UpdateWine(id uint, w *pkg.Wine) error {
	var wine *pkg.Wine
	s.DB.First(wine, id)

	s.DB.Model(wine).Updates(w)
	return nil
}
