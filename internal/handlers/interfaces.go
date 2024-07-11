package handlers

import (
	"github.com/gin-gonic/gin"
)

type Burgers interface {
	BurgersSearch(ctx *gin.Context)

	BurgerCreate(ctx *gin.Context)
	BurgerLookupByID(ctx *gin.Context)

	BurgerFindRandom(ctx *gin.Context)

	BurgersListAllStartingLetters(ctx *gin.Context)
	BurgersListAllByFirstLetter(ctx *gin.Context)
}

type Ingredients interface {
	IngredientFindByName(ctx *gin.Context)
	IngredientLookupByID(ctx *gin.Context)
	IngredientsReadAll(ctx *gin.Context)
	IngredientsSearch(ctx *gin.Context)
}
