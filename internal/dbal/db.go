package dbal

import (
	"app/internal/dbal/dao"
	"app/internal/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	db *gorm.DB
)

func Db() *gorm.DB {
	return db
}

func OpenDb() (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	dbPath := fmt.Sprintf("%s/db/burgers.sqlite", pwd)

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return
}

func CloseDb() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func AutoMigrate() error {
	return db.AutoMigrate(
		// &model.BurgerIngredient{},
		&model.Burger{},
		&model.Ingredient{},
	)
}

func NewBurgersDao() dao.BurgersDao {
	return dao.NewBurgersDao(db)
}

func NewIngredientsDao() dao.IngredientsDao {
	return dao.NewIngredientsDao(db)
}
