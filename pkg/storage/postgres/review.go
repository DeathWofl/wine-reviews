package postgres

import (
	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

type ReviewService struct {
	DB *gorm.DB
}

func (s *ReviewService) Review(id uint) (*model.Review, error) {
	var review model.Review
	s.DB.First(&review, id)
	return &review, nil
}

func (s *ReviewService) Reviews() ([]*model.Review, error) {
	var reviews []*model.Review
	s.DB.Find(&reviews)
	return reviews, nil
}

func (s *ReviewService) CreateReview(w *model.Review) (*model.Review, error) {
	s.DB.Create(w)
	return w, nil
}

func (s *ReviewService) DeleteReview(id uint) error {
	s.DB.Delete(&model.Review{}, id)
	return nil
}

func (s *ReviewService) UpdateReview(id uint, w *model.Review) error {
	var review *model.Review
	s.DB.First(&review, id)

	s.DB.Model(review).Updates(w)
	return nil
}
