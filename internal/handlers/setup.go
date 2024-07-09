package handlers

import (
	"github.com/gin-gonic/gin"
)

func Setup(myRouter *gin.Engine) {

	myRouter.Static("/static", "./front-end/static")

	// === panedrone: type "http://localhost:8080" to render index.html

	myRouter.StaticFile("/", "./front-end/index.html")

	/////////////////////////////////////////

	groupApi := myRouter.Group("/api")

	/////////////////////////////////////////

	burgersHandles := NewBurgersHandles()
	{
		burgersList := groupApi.Group("/burgers")
		burgersList.POST("/", burgersHandles.Create)
		burgersList.GET("/:burger_id", burgersHandles.Read)
		burgersList.GET("/search", burgersHandles.FindByName)
		burgersList.GET("/index", burgersHandles.ListAllStartingLetters)
		burgersList.GET("/by-first", burgersHandles.ListAllByFirstLetter)
	}
}
