package handlers

import (
	"github.com/gin-gonic/gin"
)

type Burgers interface {
	BurgersSearchByName(ctx *gin.Context)
	BurgersSearchByIngredient(ctx *gin.Context)

	BurgerCreate(ctx *gin.Context)
	BurgerLookupByID(ctx *gin.Context)

	BurgerFindRandom(ctx *gin.Context)

	BurgersListAllStartingLetters(ctx *gin.Context)
	BurgersListAllByFirstLetter(ctx *gin.Context)
}

type Ingredients interface {
	IngredientSearchByName(ctx *gin.Context)
	IngredientLookupByID(ctx *gin.Context)
}
