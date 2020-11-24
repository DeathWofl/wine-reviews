package pkg

import (
	"gorm.io/gorm"
)

type Winery struct {
	gorm.Model
	Name     string
	Location string
	Stars    *int
	Wines    []Wine
}

type WineryService interface {
	Winery(id uint) (*Winery, error)
	Winerys() *[]Winery
	CreateWinery(w *Winery) error
	DeleteWinery(id uint) error
	UpdateWinery(id uint, w *Winery) error
}
