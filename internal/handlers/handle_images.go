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

// UploadBurgerImage
//
//	@Summary	Upload Burger Image
//	@Tags		Burger Images
//	@Id			UploadBurgerImage
//	@Produce	json
//	@Success	201
//	@Failure	500
//	@Security	none
//	@Router		/images [post]
//	@Param		Content-Type: header	string	true	"Constant" default(multipart/form-data; boundary=WebAppBoundary)
//	@Param		Content-Disposition	header	string	true	"Insert image file name here" default(form-data; name="file"; filename="<Insert image file name here>")
func UploadBurgerImage(ctx *gin.Context) {
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
