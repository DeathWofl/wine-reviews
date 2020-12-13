package domain

import (
	"context"

	"github.com/deathwofl/wine-reviews/graph/model"
)

func (d *Domain) CreateWine(ctx context.Context, input model.NewWineInput) (*model.Wine, error) {
	return d.WineService.CreateWine(&model.Wine{
		Name:     input.Name,
		WineryID: uint(input.WineryID),
		ShortDes: input.ShortDes,
	})
}

func (d *Domain) UpdateWine(ctx context.Context, id int, input model.UpdateWineInput) (*model.Wine, error) {
	return d.WineService.UpdateWine(uint(id), model.Wine{
		Name:     input.Name,
		ShortDes: input.ShortDes,
	})
}

func (d *Domain) DeleteWine(ctx context.Context, id int) (bool, error) {
	return true, d.WineService.DeleteWine(uint(id))
}

func (d *Domain) Wines(ctx context.Context, filter *model.WineFilter, limit *int) ([]*model.Wine, error) {
	return d.WineService.Wines(filter, limit)
}

func (d *Domain) Wine(ctx context.Context, obj *model.Review) (*model.Wine, error) {
	return d.WineService.Wine(obj.WineID)
}
