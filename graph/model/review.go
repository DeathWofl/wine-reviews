package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Score  int    `json:"score"`
	Text   string `json:"text"`
	WineID uint   `json:"wineid"`
	UserID uint   `json:"userid"`
}

type ReviewService interface {
	Review(id uint) (*Review, error)
	Reviews() ([]*Review, error)
	CreateReview(w *Review) (*Review, error)
	DeleteReview(id uint) error
	UpdateReview(id uint, w *Review) error
}
