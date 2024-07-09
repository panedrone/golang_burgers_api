package handlers

import (
	"app/internal/dbal"
	"app/internal/dbal/dao"
	"app/internal/handlers/request"
	"app/internal/handlers/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type burgersHandles struct {
	bDao dao.BurgersDao
}

func NewBurgersHandles() BurgersHandles {
	return &burgersHandles{
		bDao: dbal.NewBurgersDao(),
	}
}

// FindByName
//
//	@Summary	Search Burgers by Name
//	@Tags		Burgers
//	@Id			Burgers_FindByName
//	@Produce	json
//	@Success	200	{object}	[]model.Burger	"Burgers list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/search [get]
//	@Param		name	query		string	false	"Name Key"	example(abc)
func (b burgersHandles) FindByName(ctx *gin.Context) {
	name := ctx.Query("name")
	res, err := b.bDao.FindByName(ctx, name)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

func (b burgersHandles) FindByIngredient(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Create
//
//	@Summary	Create Burger
//	@Tags		Burgers
//	@Id			Burgers_Create
//	@Accept		json
//	@Success	201
//	@Failure	500
//	@Security	none
//	@Router		/burgers [post]
//	@Param		json	body	model.Burger	true	"Burger data"
func (b burgersHandles) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Read
//
//	@Summary	Lookup full Burger details by id
//	@Tags		Burgers
//	@Id			Burgers_Read
//	@Produce	json
//	@Success	200	{object}	model.Burger	"Burger data"
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Security	none
//	@Router		/burgers/{burger_id} [get]
//	@Param		burger_id	path	integer	true	"Burger ID"
func (b burgersHandles) Read(ctx *gin.Context) {
	uri, err := request.BindBurgerUri(ctx)
	if err != nil {
		resp.Abort400BadUri(ctx, err)
		return
	}
	res, err := b.bDao.Read(ctx, uri.BurgerId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

func (b burgersHandles) FindRandom(ctx *gin.Context) {

	panic("implement me")
}

// ListAllStartingLetters
//
//	@Summary	List All starting Letters of Burger Names
//	@Tags		Burgers
//	@Id			Burgers_ListAllStartingLetters
//	@Produce	json
//	@Success	200	{object}	[]string	"Letter list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/index [get]
func (b burgersHandles) ListAllStartingLetters(ctx *gin.Context) {
	res, err := b.bDao.ListAllStartingLetters(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, res)
}

// ListAllByFirstLetter
//
//	@Summary	List All Burgers By First Letter
//	@Tags		Burgers
//	@Id			Burgers_ListAllByFirstLetter
//	@Produce	json
//	@Success	200	{object}	[]model.Burger	"Burgers list"
//	@Failure	500
//	@Security	none
//	@Router		/burgers/by-first [get]
//	@Param		letter	query		string	false	"First Letter of Burger Name"	example(A)
func (b burgersHandles) ListAllByFirstLetter(ctx *gin.Context) {
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
