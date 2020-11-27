package graph

import (
	"github.com/deathwofl/wine-reviews/pkg/storage/postgres"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService   postgres.UserService
	ReviewService postgres.ReviewService
	WineService   postgres.WineService
	WineryService postgres.WineryService
}
