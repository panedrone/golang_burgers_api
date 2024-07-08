package dao

import (
	"app/internal/model"
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type burgersDao struct {
	db *gorm.DB
}

func NewBurgersDao(db *gorm.DB) BurgersDao {
	return &burgersDao{db: db}
}

func (b *burgersDao) Create(ctx context.Context, item *model.Burger) (err error) {
	err = b.db.Model(item).Association(model.BurgerIngredients).Error
	if err != nil {
		return
	}
	err = b.db.WithContext(ctx).Create(item).Error
	return
}

func (b *burgersDao) Read(ctx context.Context, bID int64) (res *model.Burger, err error) {
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Where("id = ?", bID).First(&res).Error
	return
}

func (b *burgersDao) FindRandom(ctx context.Context) (*model.Burger, error) {
	//TODO implement me
	panic("implement me")
}

func (b *burgersDao) FindByName(ctx context.Context, cName string) (res []*model.Burger, err error) {
	key := fmt.Sprintf("%%%s%%", strings.ToLower(cName))
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Where("LOWER(burger_name) LIKE ?", key).Find(&res).Error
	return
}

func (b *burgersDao) FindByIngredient(ctx context.Context, cIngredientName string) (res []*model.Burger, err error) {
	key := fmt.Sprintf("%%%s%%", strings.ToLower(cIngredientName))
	subQuery := b.db.Model(&model.Ingredient{}).Select("DISTINCT(burger_ingredients.burger_id)").Joins("inner join burger_ingredients ON id = burger_ingredients.ingredient_id").
		Where("LOWER(ingredient_name) LIKE ?", key)
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Where("id in (?)", subQuery).Find(&res).Error
	return
}
