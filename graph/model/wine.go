package model

import (
	"gorm.io/gorm"
)

type Wine struct {
	gorm.Model
	Name     string    `json:"name"`
	ShortDes string    `json:"shortdes"`
	WineryID uint      `json:"wineryid"`
	Reviews  []*Review `json:"reviews"`
}

type WineService interface {
	Wine(id uint) (*Wine, error)
	Wines() ([]*Wine, error)
	CreateWine(w *Wine) (*Wine, error)
	DeleteWine(id string) error
	UpdateWine(id string, w *Wine) error
}
