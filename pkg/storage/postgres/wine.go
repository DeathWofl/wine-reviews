package postgres

import (
	"fmt"

	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

type WineService struct {
	DB *gorm.DB
}

func (s *WineService) Wine(id uint) (*model.Wine, error) {
	var wine model.Wine
	s.DB.First(&wine, id)
	return &wine, nil
}

func (s *WineService) Wines(filter *model.WineFilter, limit *int) ([]*model.Wine, error) {
	var wines []*model.Wine
	fmt.Printf("Limit: %v\n", *limit)
	query := s.DB
	if filter != nil {
		if limit != nil {
			query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name)).Limit(*limit)
		}
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	query.Find(&wines)
	return wines, nil
}

func (s *WineService) CreateWine(w *model.Wine) (*model.Wine, error) {
	s.DB.Create(w)
	return w, nil
}

func (s *WineService) DeleteWine(id uint) error {

	s.DB.Delete(&model.Wine{}, id)
	return nil
}

func (s *WineService) UpdateWine(id uint, w model.Wine) (*model.Wine, error) {
	var wine model.Wine
	s.DB.First(&wine, id)
	s.DB.Model(&wine).Updates(w)
	return &wine, nil
}
