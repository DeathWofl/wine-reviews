package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
	return r.ReviewService.Reviews()
}

func (r *reviewResolver) ID(ctx context.Context, obj *model.Review) (int, error) {
	return int(obj.ID), nil
}

func (r *reviewResolver) Wine(ctx context.Context, obj *model.Review) (*model.Wine, error) {
	return r.WineService.Wine(obj.WineID)
}

func (r *reviewResolver) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	return r.UserService.User(obj.UserID)
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) ID(ctx context.Context, obj *model.Wine) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) Winery(ctx context.Context, obj *model.Wine) (*model.Winery, error) {
	return r.WineryService.Winery(obj.WineryID)
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
