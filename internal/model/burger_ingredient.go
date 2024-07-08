package model

// !!! don't use for migrate

type BurgerIngredient struct {
	BurgerID     int64 `gorm:"column:burger_id;primaryKey;autoIncrement:false"`
	IngredientID int64 `gorm:"column:ingredient_id;primaryKey;autoIncrement:false"`
}
