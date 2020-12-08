package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/deathwofl/wine-reviews/graph/generated"
	"github.com/deathwofl/wine-reviews/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := Validation(ctx, input)

	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := Validation(ctx, input)

	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.Login(ctx, input)
}

func (r *mutationResolver) CreateWinery(ctx context.Context, input model.NewWineryInput) (*model.Winery, error) {
	return r.Domain.CreateWinery(ctx, input)
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReviewInput) (*model.Review, error) {
	return r.Domain.CreateReview(ctx, input)
}

func (r *mutationResolver) CreateWine(ctx context.Context, input model.NewWineInput) (*model.Wine, error) {
	return r.Domain.CreateWine(ctx, input)
}

func (r *mutationResolver) UpdateWine(ctx context.Context, id int, input model.UpdateWineInput) (*model.Wine, error) {
	return r.Domain.UpdateWine(ctx, id, input)
}

func (r *mutationResolver) UpdateWinery(ctx context.Context, id int, input model.UpdateWineryInput) (*model.Winery, error) {
	return r.Domain.UpdateWinery(ctx, id, input)
}

func (r *mutationResolver) UpdateReview(ctx context.Context, id int, input model.UpdateReviewInput) (*model.Review, error) {
	return r.Domain.UpdateReview(ctx, id, input)
}

func (r *mutationResolver) DeleteWine(ctx context.Context, id int) (bool, error) {
	return r.Domain.DeleteWine(ctx, id)
}

func (r *mutationResolver) DeleteWinery(ctx context.Context, id int) (bool, error) {
	return r.Domain.DeleteWinery(ctx, id)
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id int) (bool, error) {
	return r.Domain.DeleteReview(ctx, id)
}

func (r *queryResolver) Wines(ctx context.Context, filter *model.WineFilter, limit *int) ([]*model.Wine, error) {
	return r.Domain.Wines(ctx, filter, limit)
}

func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	return r.Domain.Reviews(ctx)
}

func (r *reviewResolver) ID(ctx context.Context, obj *model.Review) (int, error) {
	return int(obj.ID), nil
}

func (r *reviewResolver) Wine(ctx context.Context, obj *model.Review) (*model.Wine, error) {
	return r.Domain.Wine(ctx, obj)
}

func (r *reviewResolver) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	return r.Domain.User(ctx, obj)
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) ID(ctx context.Context, obj *model.Wine) (int, error) {
	return int(obj.ID), nil
}

func (r *wineResolver) Winery(ctx context.Context, obj *model.Wine) (*model.Winery, error) {
	return r.Domain.Winery(ctx, obj)
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
