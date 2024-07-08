package dbal

import (
	"app/internal/dbal/dao"
	"app/internal/model"
	"app/internal/shared"
	"context"
	"os"
	"path"
	"runtime"
	"testing"
)

var ctx context.Context

func TestMain(m *testing.M) {
	{
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../..")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}
	}

	err := OpenDb()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		_ = CloseDb()
	}()
	ctx = context.Background()
	code := m.Run()
	// .................... clean up
	os.Exit(code)
}

func Test_FindByName(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.FindByName(ctx, "barbiE")
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_FindByIngredient(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.FindByIngredient(ctx, "Tomato")
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_Read(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.Read(ctx, 1)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_Create(t *testing.T) {

	tx := Db().Begin()

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

	res, err := bDao.Read(ctx, dbBurger.ID)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_FindRandom(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.FindRandom(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	shared.PrintLnCyan(res)
}

func Test_ListAllByFirstLetter(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.ListAllByFirstLetter(ctx, '2')
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}

func Test_ListAllStartingLetters(t *testing.T) {
	bDao := NewBurgersDao()
	res, err := bDao.ListAllStartingLetters(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}
	for _, v := range res {
		shared.PrintLnCyan(v)
	}
}
