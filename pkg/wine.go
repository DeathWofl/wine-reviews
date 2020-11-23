package pkg

import (
	"gorm.io/gorm"
)

type Wine struct {
	gorm.Model
	Name     string
	ShortDes string
	Winery   *Winery
}

type WineService interface {
	Wine(id string) (*Wine, error)
	Wines() []*Wine
	CreateWine(w *Wine) error
	DeleteWine(id string) error
	UpdateWine(id string, w *Wine) error
}
