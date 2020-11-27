package model

type Review struct {
	ID     uint   `json:"id"`
	Score  int    `json:"score"`
	Text   string `json:"text"`
	WineID uint   `json:"wineid"`
	UserID uint   `json:"userid"`
}

type ReviewService interface {
	Review(id uint) (*Review, error)
	Reviews() []*Review
	CreateReview(w *Review) error
	DeleteReview(id uint) error
	UpdateReview(id uint, w *Review) error
}