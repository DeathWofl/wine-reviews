package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/deathwofl/wine-reviews/graph/generated"
	"github.com/deathwofl/wine-reviews/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWinery(ctx context.Context, input model.NewWinery) (*model.Winery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReview) (*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWine(ctx context.Context, input model.NewWine) (*model.Wine, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	return reviews, nil
}

func (r *reviewResolver) ID(ctx context.Context, obj *model.Review) (int, error) {
	return int(obj.ID), nil
}

func (r *reviewResolver) Wine(ctx context.Context, obj *model.Review) (*model.Wine, error) {
	for _, wine := range wines {
		if wine.ID == obj.WineID {
			return wine, nil
		}
	}
	return nil, errors.New("Wine with id not exits.")
}

func (r *reviewResolver) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	for _, user := range users {
		if user.ID == obj.UserID {
			return user, nil
		}
	}
	return nil, errors.New("User with id not exits.")
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) ID(ctx context.Context, obj *model.Wine) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) Winery(ctx context.Context, obj *model.Wine) (*model.Winery, error) {
	for _, winery := range winerys {
		if winery.ID == obj.WineryID {
			return winery, nil
		}
	}
	return nil, errors.New("Winery with id not exits.")
}

func (r *wineryResolver) ID(ctx context.Context, obj *model.Winery) (int, error) {
	return int(obj.ID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Wine returns generated.WineResolver implementation.
func (r *Resolver) Wine() generated.WineResolver { return &wineResolver{r} }

// Winery returns generated.WineryResolver implementation.
func (r *Resolver) Winery() generated.WineryResolver { return &wineryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type wineResolver struct{ *Resolver }
type wineryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var reviews = []*model.Review{
	{
		ID:     1,
		Score:  5,
		Text:   "First Desc.",
		WineID: 1,
		UserID: 1,
	},
	{
		ID:     2,
		Score:  4,
		Text:   "Second Desc.",
		WineID: 1,
		UserID: 1,
	},
	{
		ID:     3,
		Score:  3,
		Text:   "Third Desc.",
		WineID: 2,
		UserID: 1,
	},
	{
		ID:     4,
		Score:  2,
		Text:   "Four Desc.",
		WineID: 2,
		UserID: 2,
	},
	{
		ID:     5,
		Score:  1,
		Text:   "Five Desc.",
		WineID: 3,
		UserID: 2,
	},
}
var wines = []*model.Wine{
	{
		ID:       1,
		Name:     "Merlot",
		ShortDes: "Short Desc 1.",
		WineryID: 1,
	},
	{
		ID:       2,
		Name:     "Tempranillo",
		ShortDes: "Short Desc 2",
		WineryID: 1,
	},
	{
		ID:       3,
		Name:     "Merlot",
		ShortDes: "Short Desc 3.",
		WineryID: 1,
	},
}
var winerys = []*model.Winery{
	{
		ID:       1,
		Name:     "Red Winery",
		Location: "Spain",
		Stars:    5,
	},
}
var users = []*model.User{
	{
		ID:       1,
		Username: "username1",
		Password: "pass123",
		Email:    "username1@example.com",
	},
	{
		ID:       2,
		Username: "username2",
		Password: "pass123",
		Email:    "username2@example.com",
	},
}
