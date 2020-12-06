package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/deathwofl/wine-reviews/graph/generated"
	"github.com/deathwofl/wine-reviews/graph/model"
	"github.com/deathwofl/wine-reviews/pkg/middleware"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	confirmEmail, err := r.UserService.UserbyEmail(input.Email)
	if confirmEmail.Email != "" {
		return nil, errors.New("email already in use")
	}

	confirmUsername, err := r.UserService.UserbyUsername(input.Username)
	if confirmUsername.Username != "" {
		return nil, errors.New("username already in use")
	}

	user := &model.User{
		Username:  input.Username,
		Lastname:  input.LastName,
		FirstName: input.FirstName,
		Email:     input.Email,
	}

	err = middleware.HashPassword(input.Password, user)
	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return nil, errors.New("Error hashing password")
	}

	_, err = r.UserService.CreateUser(user)
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return nil, errors.New("Error creating user")
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return nil, errors.New("Error generating token")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UserService.UserbyEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password are wrong")
	}

	err = middleware.ComparePassword(input.Password, user)
	if err != nil {
		log.Printf("Error while comparing password: %v", err)
		return nil, errors.New("email or password are wrong 2")
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateWinery(ctx context.Context, input model.NewWineryInput) (*model.Winery, error) {
	return r.WineryService.CreateWinery(&model.Winery{
		Name:     input.Name,
		Location: input.Location,
		Stars:    input.Stars,
	})
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReviewInput) (*model.Review, error) {
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, errors.New("Unathorizated")
	}
	return r.ReviewService.CreateReview(&model.Review{
		UserID: currentUser.ID,
		Score:  input.Score,
		Text:   input.Text,
		WineID: uint(input.WineID),
	})
}

func (r *mutationResolver) CreateWine(ctx context.Context, input model.NewWineInput) (*model.Wine, error) {
	return r.WineService.CreateWine(&model.Wine{
		Name:     input.Name,
		WineryID: uint(input.WineryID),
		ShortDes: input.ShortDes,
	})
}

func (r *mutationResolver) UpdateWine(ctx context.Context, id int, input model.UpdateWineInput) (*model.Wine, error) {
	return r.WineService.UpdateWine(uint(id), model.Wine{
		Name:     input.Name,
		ShortDes: input.ShortDes,
	})
}

func (r *mutationResolver) UpdateWinery(ctx context.Context, id int, input model.UpdateWineryInput) (*model.Winery, error) {
	return r.WineryService.UpdateWinery(uint(id), model.Winery{
		Name:     input.Name,
		Location: input.Location,
		Stars:    *input.Stars,
	})
}

func (r *mutationResolver) UpdateReview(ctx context.Context, id int, input model.UpdateReviewInput) (*model.Review, error) {
	return r.ReviewService.UpdateReview(uint(id), model.Review{
		Score: input.Score,
		Text:  input.Text,
	})
}

func (r *mutationResolver) DeleteWine(ctx context.Context, id int) (bool, error) {
	return true, r.WineService.DeleteWine(uint(id))
}

func (r *mutationResolver) DeleteWinery(ctx context.Context, id int) (bool, error) {
	return true, r.WineryService.DeleteWinery(uint(id))
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id int) (bool, error) {
	return true, r.ReviewService.DeleteReview(uint(id))
}

func (r *queryResolver) Wines(ctx context.Context, filter *model.WineFilter, limit *int) ([]*model.Wine, error) {
	return r.WineService.Wines(filter, limit)
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
