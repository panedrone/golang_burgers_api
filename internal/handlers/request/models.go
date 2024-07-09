package request

import (
	"app/internal/handlers/resp"
	"github.com/gin-gonic/gin"
)

type BurgerUri struct {
	BurgerId int64 `uri:"burger_id"  binding:"required" example:"1"`
}

func BindBurgerUri(ctx *gin.Context) (*BurgerUri, error) {
	var uri BurgerUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}
