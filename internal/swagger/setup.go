package swagger

import (
	_ "app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(myRouter *gin.Engine) {
	myRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
