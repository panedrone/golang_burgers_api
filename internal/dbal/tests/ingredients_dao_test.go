package dbal

import (
	"app/internal/dbal"
	"app/internal/shared"
	"testing"
)

func Test_Ingredients_FindByName(t *testing.T) {
	iDao := dbal.NewIngredientsDao()
	res, err := iDao.FindByName(ctx, "tomatoes")
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_Ingredients_LookupIngredientByID(t *testing.T) {
	iDao := dbal.NewIngredientsDao()
	res, err := iDao.LookupIngredientByID(ctx, 1)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}
