package dao

import (
	"app/internal/model"
	"context"
)

type BurgersDao interface {
	FindByName(ctx context.Context, cName string) (res []*model.Burger, err error)
	FindByIngredient(ctx context.Context, cIngredientName string) (res []*model.Burger, err error)

	Create(ctx context.Context, item *model.Burger) error
	Read(ctx context.Context, cID int64) (*model.Burger, error)

	FindRandom(ctx context.Context) (*model.Burger, error)
}
