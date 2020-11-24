package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/deathwofl/wine-reviews/graph/generated"
	"github.com/deathwofl/wine-reviews/graph/model"
	"github.com/deathwofl/wine-reviews/pkg"
	storage "github.com/deathwofl/wine-reviews/pkg/storage/mysql"
)

var Users storage.UserService
var Wines storage.WineService
var Winerys storage.WineryService
var Reviews storage.ReviewService

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWinery(ctx context.Context, input model.NewWinery) (*model.Winery, error) {
	var winery pkg.Winery
	winery.Name = input.Name
	winery.Location = input.Location
	winery.Stars = input.Stars

	Winerys.CreateWinery(&winery)

	return &model.Winery{
		Name:     winery.Name,
		Location: winery.Location,
		Stars:    winery.Stars,
	}, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReview) (*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWine(ctx context.Context, input model.NewWine) (*model.Wine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	var reviews []*model.Review
	var star *int
	star = new(int)
	*star = 5
	basicReview := model.Review{
		Wine: &model.Wine{
			Name:     "Bondue",
			ShortDes: "Basic wine, example.",
			Winery: &model.Winery{
				Name:     "Le Caset White",
				Location: "Spain, Costuella",
				Stars:    star,
			},
		},
		Text:  "Basic review, example.",
		Score: 5,
		User: &model.User{
			Username: "exampleuser",
			Password: "examplePassword",
			Email:    "example@example.com",
		},
	}

	reviews = append(reviews, &basicReview)

	return reviews, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
