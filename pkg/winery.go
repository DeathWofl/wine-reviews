package pkg

type Winery struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Starts   *int   `json:"starts"`
}

type WineryService interface {
	Winery(id string) (*Winery, error)
	Winerys() []*Winery
	CreateWinery(w *Winery) error
	DeleteWinery(id string) error
	UpdateWinery(id string, w *Winery) error
}
