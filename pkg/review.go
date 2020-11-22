package pkg

type Review struct {
	ID    uint
	Wine  *Wine
	User  *User
	Score int
	Text  string
}

type ReviewService interface {
	Review(id string) (*Review, error)
	Reviews() []*Review
	CreateReview(w *Review) error
	DeleteReview(id string) error
	UpdateReview(id string, w *Review) error
}
