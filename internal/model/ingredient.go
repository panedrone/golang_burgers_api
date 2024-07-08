package model

type Ingredient struct {
	ID   int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name string `json:"name" gorm:"column:ingredient_name;size:128;uniqueIndex;"`
}
