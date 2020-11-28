package model

type Wine struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	ShortDes string    `json:"shortdes"`
	WineryID uint      `json:"wineryid"`
	Reviews  []*Review `json:"reviews"`
}

type WineService interface {
	Wine(id string) (*Wine, error)
	Wines() ([]*Wine, error)
	CreateWine(w *Wine) (*Wine, error)
	DeleteWine(id string) error
	UpdateWine(id string, w *Wine) error
}
