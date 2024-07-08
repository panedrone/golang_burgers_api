package handlers

import (
	"github.com/gin-gonic/gin"
)

type BurgersHandles interface {
	FindByName(ctx *gin.Context)
	FindByIngredient(ctx *gin.Context)

	Create(ctx *gin.Context)
	Read(ctx *gin.Context)

	FindRandom(ctx *gin.Context)

	ListAllStartingLetters(ctx *gin.Context)
	ListAllByFirstLetter(ctx *gin.Context)
}
