package handlers

import (
	"app/internal/handlers/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	UrlPath = "static/images"

	ImgFolder = "front-end/" + UrlPath
)

func PostFile(ctx *gin.Context) {
	// single file
	file, err := ctx.FormFile("file")
	if err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
		return
	}

	// log.Println(file.Filename)

	dst := fmt.Sprintf("%s/%s", ImgFolder, file.Filename)

	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}

	url := fmt.Sprintf("%s/%s", UrlPath, file.Filename)
	ctx.String(http.StatusOK, url)
}
