package model

import (
	"gorm.io/gorm"
)

type Winery struct {
	gorm.Model
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Stars    int     `json:"stars"`
	Wines    []*Wine `json:"wines"`
}

type WineryService interface {
	Winery(id uint) (*Winery, error)
	Winerys() ([]*Winery, error)
	CreateWinery(w *Winery) (*Winery, error)
	DeleteWinery(id uint) error
	UpdateWinery(id uint, w *Winery) error
}
