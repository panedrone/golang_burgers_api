package handlers

import (
	"app/internal/dbal"
	"app/internal/dbal/dao"
	"app/internal/handlers/request"
	"app/internal/handlers/resp"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ingredients struct {
	iDao dao.IngredientsDao
}

func NewIngredients() Ingredients {
	return &ingredients{
		iDao: dbal.NewIngredientsDao(),
	}
}

// IngredientSearchByName
//
//	@Summary	Search Ingredient by Name
//	@Tags		Ingredients
//	@Id			IngredientSearchByName
//	@Produce	json
//	@Success	200	{object}	model.Ingredient	"Ingredient"
//	@Failure	500
//	@Security	none
//	@Router		/ingredients/search [get]
//	@Param		name	query		string	false	"Name Key"	example(abc)
func (i *ingredients) IngredientSearchByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if len(name) == 0 {
		resp.Abort400hBadRequest(ctx, "name is required")
		return
	}
	d := dbal.NewIngredientsDao()
	var res *model.Ingredient // === panedrone: for "swag init"
	res, err := d.FindByName(ctx, name)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// IngredientLookupByID
//
//	@Summary	Lookup full Ingredient details by id
//	@Tags		Ingredients
//	@Id			IngredientLookupByID
//	@Produce	json
//	@Success	200	{object}	model.Ingredient	"Ingredient data"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/ingredients/{ingredient_Id} [get]
//	@Param		ingredient_Id	path	integer	true	"Ingredient ID"
func (i *ingredients) IngredientLookupByID(ctx *gin.Context) {
	uri, err := request.BindIngredientUri(ctx)
	if err != nil {
		resp.Abort400BadUri(ctx, err)
		return
	}
	d := dbal.NewIngredientsDao()
	var res *model.Ingredient // === panedrone: for "swag init"
	res, err = d.LookupIngredientByID(ctx, uri.IngredientId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)

}

// IngredientsReadAll
//
//	@Summary	Lookup all Ingredients
//	@Tags		Ingredients
//	@Id			IngredientsReadAll
//	@Produce	json
//	@Success	200	{object}	[]model.Ingredient	"Ingredient List"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/ingredients [get]
func (i *ingredients) IngredientsReadAll(ctx *gin.Context) {
	d := dbal.NewIngredientsDao()
	res, err := d.ReadAll(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}
