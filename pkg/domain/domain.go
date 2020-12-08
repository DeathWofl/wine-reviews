package domain

import (
	"github.com/deathwofl/wine-reviews/pkg/storage/postgres"
)

type Domain struct {
	WineService   postgres.WineService
	WineryService postgres.WineryService
	ReviewService postgres.ReviewService
	UserService   postgres.UserService
}

func New(wineService postgres.WineService, wineryService postgres.WineryService, userService postgres.UserService, reviewService postgres.ReviewService) *Domain {
	return &Domain{
		WineService:   wineService,
		WineryService: wineryService,
		ReviewService: reviewService,
		UserService:   userService,
	}
}
