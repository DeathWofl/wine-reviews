package model

type Winery struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Stars    int     `json:"stars"`
	Wines    []*Wine `json:"wines"`
}

type WineryService interface {
	Winery(id uint) (*Winery, error)
	Winerys() *[]Winery
	CreateWinery(w *Winery) error
	DeleteWinery(id uint) error
	UpdateWinery(id uint, w *Winery) error
}
