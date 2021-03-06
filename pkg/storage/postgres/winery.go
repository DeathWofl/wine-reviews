package postgres

import (
	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

type WineryService struct {
	DB *gorm.DB
}

func (s *WineryService) Winery(id uint) (*model.Winery, error) {

	var winery model.Winery
	s.DB.First(&winery, id)
	return &winery, nil
}

func (s *WineryService) Winerys() ([]*model.Winery, error) {
	var winerys []*model.Winery
	s.DB.Find(&winerys)
	return winerys, nil
}

func (s *WineryService) CreateWinery(w *model.Winery) (*model.Winery, error) {
	s.DB.Create(w)
	return w, nil
}

func (s *WineryService) DeleteWinery(id uint) error {
	s.DB.Delete(&model.Winery{}, id)
	return nil
}

func (s *WineryService) UpdateWinery(id uint, w model.Winery) (*model.Winery, error) {
	var winery model.Winery
	s.DB.First(&winery, id)
	s.DB.Model(&winery).Updates(w)
	return &winery, nil
}
