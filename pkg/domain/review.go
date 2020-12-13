package domain

import (
	"context"
	"errors"

	"github.com/deathwofl/wine-reviews/graph/model"
	"github.com/deathwofl/wine-reviews/pkg/middleware"
)

func (d *Domain) CreateReview(ctx context.Context, input model.NewReviewInput) (*model.Review, error) {
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, errors.New("Unathorizated")
	}
	return d.ReviewService.CreateReview(&model.Review{
		UserID: currentUser.ID,
		Score:  input.Score,
		Text:   input.Text,
		WineID: uint(input.WineID),
	})
}

func (d *Domain) UpdateReview(ctx context.Context, id int, input model.UpdateReviewInput) (*model.Review, error) {
	return d.ReviewService.UpdateReview(uint(id), model.Review{
		Score: input.Score,
		Text:  input.Text,
	})
}

func (d *Domain) DeleteReview(ctx context.Context, id int) (bool, error) {
	return true, d.ReviewService.DeleteReview(uint(id))
}

func (d *Domain) Reviews(ctx context.Context) ([]*model.Review, error) {
	return d.ReviewService.Reviews()
}
