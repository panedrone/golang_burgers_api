package dao

import (
	"app/internal/model"
	"context"
)

type BurgersDao interface {
	SearchByName(ctx context.Context, cName string) (res []*model.Burger, err error)
	SearchByIngredient(ctx context.Context, cIngredientName string) (res []*model.Burger, err error)

	Create(ctx context.Context, item *model.Burger) error
	LookupByID(ctx context.Context, cID int64) (*model.Burger, error)

	FindRandom(ctx context.Context) (*model.Burger, error)

	ListAllStartingLetters(ctx context.Context) (res []string, err error)
	ListAllByFirstLetter(ctx context.Context, letter rune) (res []*model.Burger, err error)
}

type IngredientsDao interface {
	FindByName(ctx context.Context, cName string) (res *model.Ingredient, err error)
	LookupIngredientByID(ctx context.Context, iID int64) (*model.Ingredient, error)
}
