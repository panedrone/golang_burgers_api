package dbal

import (
	"app/internal/dbal"
	"app/internal/dbal/dao"
	"app/internal/model"
	"app/internal/shared"
	"context"
	"testing"
)

func Test_Burgers_SearchByName(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.SearchByName(ctx, "barbiE")
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_Burgers_SearchByIngredient(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.SearchByIngredient(ctx, "Tomato")
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_Burgers_LookupByID(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.LookupByID(ctx, 1)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_Burgers_Create(t *testing.T) {

	tx := dbal.Db().Begin()

	defer func() {
		tx.Rollback()
	}()

	bDao := dao.NewBurgersDao(tx)

	dbBurger := &model.Burger{
		Name:        "burger.Name",
		Description: "burger.Description",
		Ingredients: []model.Ingredient{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
			{
				ID: 3,
			},
		},
	}

	err := bDao.Create(context.Background(), dbBurger)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}

	res, err := bDao.LookupByID(ctx, dbBurger.ID)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_Burgers_FindRandom(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.FindRandom(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_Burgers_ListAllByFirstLetter(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.ListAllByFirstLetter(ctx, '2')
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_Burgers_ListAllStartingLetters(t *testing.T) {
	bDao := dbal.NewBurgersDao()
	res, err := bDao.ListAllStartingLetters(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}
