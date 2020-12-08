package domain

import (
	"context"

	"github.com/deathwofl/wine-reviews/graph/model"
)

func (d *Domain) CreateWinery(ctx context.Context, input model.NewWineryInput) (*model.Winery, error) {
	return d.WineryService.CreateWinery(&model.Winery{
		Name:     input.Name,
		Location: input.Location,
		Stars:    input.Stars,
	})
}

func (d *Domain) UpdateWinery(ctx context.Context, id int, input model.UpdateWineryInput) (*model.Winery, error) {
	return d.WineryService.UpdateWinery(uint(id), model.Winery{
		Name:     input.Name,
		Location: input.Location,
		Stars:    *input.Stars,
	})
}

func (d *Domain) DeleteWinery(ctx context.Context, id int) (bool, error) {
	return true, d.WineryService.DeleteWinery(uint(id))
}

func (d *Domain) Winery(ctx context.Context, obj *model.Wine) (*model.Winery, error) {
	return d.WineryService.Winery(obj.WineryID)
}
