package mysql

import (
	"github.com/deathwofl/wine-reviews/pkg"
	"gorm.io/gorm"
)

type ReviewService struct {
	DB *gorm.DB
}

func (s *ReviewService) Review(id uint) (*pkg.Review, error) {

	var review *pkg.Review
	s.DB.First(review, id)
	return review, nil
}

func (s *ReviewService) Reviews() []*pkg.Review {
	var reviews []*pkg.Review

	s.DB.Find(reviews)
	return reviews
}

func (s *ReviewService) CreateReview(w *pkg.Review) error {

	s.DB.Create(w)
	return nil
}

func (s *ReviewService) DeleteReview(id uint) error {

	s.DB.Delete(&pkg.Review{}, id)
	return nil
}

func (s *ReviewService) UpdateReview(id uint, w *pkg.Review) error {
	var review *pkg.Review
	s.DB.First(review, id)

	s.DB.Model(review).Updates(w)
	return nil
}
