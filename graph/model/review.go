package model

type Review struct {
	ID     string `json:"id"`
	Score  int    `json:"score"`
	Text   string `json:"text"`
	WineID string `json:"wineid"`
	UserID string `json:"userid"`
}

// type ReviewService interface {
// 	Review(id string) (*Review, error)
// 	Reviews() []*Review
// 	CreateReview(w *Review) error
// 	DeleteReview(id string) error
// 	UpdateReview(id string, w *Review) error
// }
