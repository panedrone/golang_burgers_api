package dao

import (
	"app/internal/model"
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand/v2"
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

func (b *burgersDao) LookupByID(ctx context.Context, bID int64) (res *model.Burger, err error) {
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Where("id = ?", bID).First(&res).Error
	return
}

func (b *burgersDao) Search(ctx context.Context, cName string, cIngredientName string) (res []*model.Burger, err error) {
	if len(cIngredientName) < 2 && len(cName) < 2 {
		return
	}
	tx := b.db.WithContext(ctx).Model(&model.Burger{}).Preload(model.BurgerIngredients)
	if len(cIngredientName) > 0 {
		tx.Scopes(func(db *gorm.DB) *gorm.DB {
			key := fmt.Sprintf("%%%s%%", strings.ToLower(cIngredientName))
			subQuery := b.db.Model(&model.Ingredient{}).Select("DISTINCT(burger_ingredients.burger_id)").Joins("inner join burger_ingredients ON id = burger_ingredients.ingredient_id").
				Where("LOWER(ingredient_name) LIKE ?", key)
			return db.Where("id in (?)", subQuery).Find(&res)
		})
	}
	if len(cName) > 0 {
		tx.Scopes(func(db *gorm.DB) *gorm.DB {
			key := fmt.Sprintf("%%%s%%", strings.ToLower(cName))
			return db.Where("LOWER(burger_name) LIKE ?", key)
		})
	}
	err = tx.Find(&res).Error
	return
}

func (b *burgersDao) ListAllStartingLetters(ctx context.Context) (res []string, err error) {
	// The first character has an index of 1.
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Select("DISTINCT(UPPER(substr(burger_name, 1, 1)))").Find(&res).Error
	return
}

func (b *burgersDao) ListAllByFirstLetter(ctx context.Context, letter rune) (res []*model.Burger, err error) {
	cName := strings.ToLower(string(letter))
	key := fmt.Sprintf("%s%%", cName)
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Where("LOWER(burger_name) LIKE ?", key).Find(&res).Error
	return
}

func (b *burgersDao) FindRandom(ctx context.Context) (res *model.Burger, err error) {
	var count int64
	err = b.db.WithContext(ctx).Model(&model.Burger{}).Count(&count).Error
	if err != nil {
		return
	}
	offset := rand.IntN(int(count))
	err = b.db.WithContext(ctx).Model(&model.Burger{}).
		Preload(model.BurgerIngredients).
		Offset(offset).First(&res).Error
	return
}
