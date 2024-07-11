package handlers

import "github.com/gin-gonic/gin"

func SetupHandlers(myRouter *gin.Engine) {

	myRouter.Static("/static", "./front-end/static")

	// === panedrone: "http://localhost:8080" to render index.html

	myRouter.StaticFile("/", "./front-end/index.html")

	/////////////////////////////////////////

	groupApi := myRouter.Group("/api")

	/////////////////////////////////////////

	{
		handleBurgers := NewBurgers()
		groupBurgers := groupApi.Group("/burgers")
		groupBurgers.POST("/", handleBurgers.BurgerCreate)
		groupBurgers.GET("/:burger_id", handleBurgers.BurgerLookupByID)
		groupBurgers.GET("/search", handleBurgers.BurgersSearch)
		groupBurgers.GET("/search/first-letters", handleBurgers.BurgersListAllStartingLetters)
		groupBurgers.GET("/search/first-letter", handleBurgers.BurgersListAllByFirstLetter)
		groupBurgers.GET("/random", handleBurgers.BurgerFindRandom)
	}
	{
		handleIngredients := NewIngredients()
		groupIngredients := groupApi.Group("/ingredients")
		groupIngredients.GET("/search", handleIngredients.IngredientSearchByName)
		groupIngredients.GET("/:ingredient_id", handleIngredients.IngredientLookupByID)
		groupIngredients.GET("/", handleIngredients.IngredientsReadAll)
	}
	{
		groupImages := groupApi.Group("/images")
		groupImages.POST("/", UploadBurgerImage)
	}
}
