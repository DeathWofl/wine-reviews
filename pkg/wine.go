package pkg

type Wine struct {
	ID       uint
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
