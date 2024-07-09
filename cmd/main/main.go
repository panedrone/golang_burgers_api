package main

import (
	"app/internal/dbal"
	"app/internal/handlers"
	"app/internal/swagger"
	"github.com/gin-gonic/gin"
	"log"
)

// @schemes	test_http
// @produce	json
// @version	0.0.1
//
// @title		Burgers API
// @in			header
// @BasePath	/api
// @accept		json
// @host		127.0.0.1:8080
func main() {
	err := dbal.OpenDb()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = dbal.CloseDb()
	}()

	gin.SetMode(gin.ReleaseMode)

	// myRouter := gin.Default()
	myRouter := gin.New()

	swagger.Setup(myRouter)

	setupHandlers(myRouter)

	log.Fatal(myRouter.Run(":8080"))
}

func setupHandlers(myRouter *gin.Engine) {

	myRouter.Static("/static", "./front-end/static")

	// === panedrone: "http://localhost:8080" to render index.html

	myRouter.StaticFile("/", "./front-end/index.html")

	/////////////////////////////////////////

	groupApi := myRouter.Group("/api")

	/////////////////////////////////////////

	{
		handleBurgers := handlers.NewBurgers()
		groupBurgers := groupApi.Group("/burgers")
		groupBurgers.POST("/", handleBurgers.BurgerCreate)
		groupBurgers.GET("/:burger_id", handleBurgers.BurgerLookupByID)
		groupBurgers.GET("/search", handleBurgers.BurgersSearchByName)
		groupBurgers.GET("/search/by-ingredient", handleBurgers.BurgersSearchByIngredient)
		groupBurgers.GET("/indexes", handleBurgers.BurgersListAllStartingLetters)
		groupBurgers.GET("/by-index", handleBurgers.BurgersListAllByFirstLetter)
		groupBurgers.GET("/random", handleBurgers.BurgerFindRandom)
	}
	{
		handleIngredients := handlers.NewIngredients()
		groupIngredients := groupApi.Group("/ingredients")
		groupIngredients.GET("/search", handleIngredients.IngredientSearchByName)
		groupIngredients.GET("/:ingredient_id", handleIngredients.IngredientLookupByID)
	}
}
