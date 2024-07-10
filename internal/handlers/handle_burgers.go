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

type burgers struct {
	bDao dao.BurgersDao
}

func NewBurgers() Burgers {
	return &burgers{
		bDao: dbal.NewBurgersDao(),
	}
}

// BurgersSearch
//
//	@Summary	Search Burgers by Name and/or Ingredient
//	@Tags		Burgers
//	@Id			BurgersSearch
//	@Produce	json
//	@Success	200	{object}	[]model.Burger	"Burgers list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/search [get]
//	@Param		b	query		string	false	"Burger Name Key"	example(abc)
//	@Param		i	query		string	false	"Ingredient Name Key"	example(abc)
func (b *burgers) BurgersSearch(ctx *gin.Context) {
	bName := ctx.Query("b")
	iName := ctx.Query("i")
	if len(bName) < 2 && len(iName) < 2 {
		resp.Abort400hBadRequest(ctx, "name is required")
		return
	}
	res, err := b.bDao.Search(ctx, bName, iName)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// BurgerCreate
//
//	@Summary	Create Burger
//	@Tags		Burgers
//	@Id			BurgerCreate
//	@Accept		json
//	@Success	201
//	@Failure	500
//	@Security	none
//	@Router		/burgers [post]
//	@Param		json	body	model.Burger	true	"Burger data"
func (b *burgers) BurgerCreate(ctx *gin.Context) {
	var req model.Burger
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	err := b.bDao.Create(ctx, &req)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
}

// BurgerLookupByID
//
//	@Summary	Lookup full Burger details by id
//	@Tags		Burgers
//	@Id			BurgerLookupByID
//	@Produce	json
//	@Success	200	{object}	model.Burger	"Burger data"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/burgers/{burger_id} [get]
//	@Param		burger_id	path	integer	true	"Burger ID"
func (b *burgers) BurgerLookupByID(ctx *gin.Context) {
	uri, err := request.BindBurgerUri(ctx)
	if err != nil {
		resp.Abort400BadUri(ctx, err)
		return
	}
	res, err := b.bDao.LookupByID(ctx, uri.BurgerId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// BurgerFindRandom
//
//	@Summary	Lookup a random Burger
//	@Tags		Burgers
//	@Id			BurgerFindRandom
//	@Produce	json
//	@Success	200	{object}	model.Burger	"Burger"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/random [get]
func (b *burgers) BurgerFindRandom(ctx *gin.Context) {
	res, err := b.bDao.FindRandom(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// BurgersListAllStartingLetters
//
//	@Summary	List All starting Letters of Burger Names
//	@Tags		Burgers
//	@Id			BurgersListAllStartingLetters
//	@Produce	json
//	@Success	200	{object}	[]string	"Letter list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/search/first-letters [get]
func (b *burgers) BurgersListAllStartingLetters(ctx *gin.Context) {
	res, err := b.bDao.ListAllStartingLetters(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// BurgersListAllByFirstLetter
//
//	@Summary	List All Burgers By First Letter
//	@Tags		Burgers
//	@Id			BurgersListAllByFirstLetter
//	@Produce	json
//	@Success	200	{object}	[]model.Burger	"Burgers list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/search/first-letter [get]
//	@Param		letter	query		string	false	"First Letter of Burger Name"	example(A)
func (b *burgers) BurgersListAllByFirstLetter(ctx *gin.Context) {
	letter := ctx.Query("letter")
	if len(letter) == 0 {
		resp.Abort400hBadRequest(ctx, "invalid query params")
		return
	}
	r := []rune(letter)[0]
	res, err := b.bDao.ListAllByFirstLetter(ctx, r)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}
