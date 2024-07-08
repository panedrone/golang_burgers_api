package model

const BurgerIngredients = "Ingredients"

type Burger struct {
	ID          int64   `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name        string  `json:"name" gorm:"column:burger_name;size:256;;uniqueIndex;not null"`
	Description string  `json:"description" gorm:"column:burger_description;size:1024;not null"`
	ImageURL    *string `json:"image_url" gorm:"column:burger_image_url;size:512"` // nullable

	Ingredients []Ingredient `json:"ingredients" gorm:"many2many:burger_ingredients"`
}
