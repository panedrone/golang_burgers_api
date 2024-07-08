package main

import (
	"app/internal/dbal"
	"app/internal/handlers"
	"app/internal/swagger"
	"github.com/gin-gonic/gin"
	"log"
)

// @schemes	http
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

	handlers.Setup(myRouter)

	log.Fatal(myRouter.Run(":8080"))
}
