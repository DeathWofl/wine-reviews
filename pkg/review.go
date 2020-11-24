package pkg

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	WineID uint
	UserID uint
	Score  int
	Text   string
}

type ReviewService interface {
	Review(id string) (*Review, error)
	Reviews() []*Review
	CreateReview(w *Review) error
	DeleteReview(id string) error
	UpdateReview(id string, w *Review) error
}
