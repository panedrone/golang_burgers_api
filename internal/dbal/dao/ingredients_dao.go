package dao

import (
	"app/internal/model"
	"context"
	"gorm.io/gorm"
)

type ingredientsDao struct {
	db *gorm.DB
}

func NewIngredientsDao(db *gorm.DB) IngredientsDao {
	return &ingredientsDao{db: db}
}

func (d *ingredientsDao) FindByName(ctx context.Context, name string) (res *model.Ingredient, err error) {
	err = d.db.WithContext(ctx).Model(&model.Ingredient{}).
		Where("LOWER(ingredient_name) = ?", name).First(&res).Error
	return
}

func (d *ingredientsDao) LookupIngredientByID(ctx context.Context, iID int64) (res *model.Ingredient, err error) {
	err = d.db.WithContext(ctx).Model(&model.Ingredient{}).
		Where("id = ?", iID).First(&res).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	err = nil
	//	return
	//}
	return
}

func (d *ingredientsDao) ReadAll(ctx context.Context) (res []*model.Ingredient, err error) {
	err = d.db.WithContext(ctx).Model(&model.Ingredient{}).Order("ingredient_name").Find(&res).Error
	return
}
