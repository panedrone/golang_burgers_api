package handlers

import (
	"app/internal/dbal"
	"app/internal/handlers/request"
	"app/internal/handlers/resp"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ingredients struct {
}

func NewIngredients() Ingredients {
	return &ingredients{}
}

// IngredientFindByName
//
//	@Summary	Search Ingredient by Name
//	@Tags		Ingredients
//	@Id			IngredientFindByName
//	@Produce	json
//	@Success	200	{object}	model.Ingredient	"Ingredient"
//	@Failure	500
//	@Security	none
//	@Router		/ingredients/find [get]
//	@Param		name	query		string	false	"Name Key"	example(abc)
func (i *ingredients) IngredientFindByName(ctx *gin.Context) {
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

// IngredientsSearch
//
//	@Summary	Search Burgers by Name and/or Ingredient
//	@Tags		Ingredients
//	@Id			IngredientsSearch
//	@Produce	json
//	@Success	200	{object}	[]model.Ingredient	"Ingredient list"
//	@Failure	500
//	@Security	none
//	@Router		/ingredients/search [get]
//	@Param		key	query		string	false	"Ingredient Name Key"	example(abc)
func (i *ingredients) IngredientsSearch(ctx *gin.Context) {
	iName := ctx.Query("key")
	if len(iName) < 1 {
		resp.Abort400hBadRequest(ctx, "invalid query params")
		return
	}
	d := dbal.NewIngredientsDao()
	res, err := d.Search(ctx, iName)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}
