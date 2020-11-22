package pkg

type Review struct {
	ID    string `json:"id"`
	Wine  *Wine  `json:"wine"`
	User  *User  `json:"user"`
	Score int    `json:"score"`
	Text  string `json:"text"`
}

type ReviewService interface {
	Review(id string) (*Review, error)
	Reviews() []*Review
	CreateReview(w *Review) error
	DeleteReview(id string) error
	UpdateReview(id string, w *Review) error
}
