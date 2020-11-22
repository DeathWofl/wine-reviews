package pkg

type Wine struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	ShortDes string  `json:"shortDes"`
	Winery   *Winery `json:"Winery"`
}

type WineService interface {
	Wine(id string) (*Wine, error)
	Wines() []*Wine
	CreateWine(w *Wine) error
	DeleteWine(id string) error
	UpdateWine(id string, w *Wine) error
}
