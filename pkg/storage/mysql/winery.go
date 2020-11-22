package mysql

import (
	"github.com/deathwofl/wine-reviews/pkg"
	"gorm.io/gorm"
)

type WineryService struct {
	DB *gorm.DB
}

func (s *WineryService) Winery(id uint) (*pkg.Winery, error) {

	var winery *pkg.Winery
	s.DB.First(winery, id)
	return winery, nil
}

func (s *WineryService) Winerys() []*pkg.Winery {
	var Winerys []*pkg.Winery

	s.DB.Find(Winerys)
	return Winerys
}

func (s *WineryService) CreateWinery(w *pkg.Winery) error {

	s.DB.Create(w)
	return nil
}

func (s *WineryService) DeleteWinery(id uint) error {

	s.DB.Delete(&pkg.Winery{}, id)
	return nil
}

func (s *WineryService) UpdateWinery(id uint, w *pkg.Winery) error {
	var winery *pkg.Winery
	s.DB.First(winery, id)

	s.DB.Model(winery).Updates(w)
	return nil
}
