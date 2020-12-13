package domain

import (
	"context"

	"github.com/deathwofl/wine-reviews/graph/model"
)

func (d *Domain) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	return d.UserService.User(obj.UserID)
}
