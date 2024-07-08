package handlers

import (
	"app/internal/dbal"
	"app/internal/dbal/dao"
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
//	@Param		json	body	request.Target	true	"Burger data"
func (b burgersHandles) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b burgersHandles) Read(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
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
//	@Router		/burgers/by-letter [get]
//	@Param		letter	query		string	false	"Firs Letter of Burger Name"	example(A)
func (b burgersHandles) ListAllByFirstLetter(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//// ReadAllByMission
////
////	@Summary	find all Burgers By Mission
////	@Tags		Burgers
////	@Id			BurgersReadAllByMission
////	@Produce	json
////	@Success	200	{object}	[]model.Target	"Burgers list"
////	@Failure	500
////	@Security	none
////	@Router		/missions/{m_id}/Burgers [get]
////	@Param		m_id	path	integer	true	"Mission ID"
//func (h *burgersHandles) ReadAllByMission(ctx *gin.Context) {
//	uri, err := request.BindMissionUri(ctx)
//	if err != nil {
//		return
//	}
//	res, err := h.dao.ReadAllByMission(ctx, uri.MId)
//	if err != nil {
//		resp.Abort500(ctx, err)
//		return
//	}
//	resp.JSON(ctx, http.StatusOK, res)
//}
//
//// ReadAllBySpyCat
////
////	@Summary	Assign Spy Cat to Target
////	@Tags		Burgers
////	@Id			BurgersReadAllBySpyCat
////	@Tags		Burgers
////	@Produce	json
////	@Success	200	{object}	[]model.Target	"Unassigned Burgers list"
////	@Failure	500
////	@Security	none
////	@Router		/spy-cats/{c_id}/Burgers [get]
////	@Param		c_id	path	integer	true	"Spy Cat ID"
//func (h *burgersHandles) ReadAllBySpyCat(ctx *gin.Context) {
//	uri, err := request.BindSpyCatUri(ctx)
//	if err != nil {
//		return
//	}
//	res, err := h.dao.ReadAllBySpyCat(ctx, uri.CId)
//	if err != nil {
//		resp.Abort500(ctx, err)
//	}
//	resp.JSON(ctx, http.StatusOK, res)
//}
//
//// Create
////
////	@Summary	create Burgers
////	@Tags		Burgers
////	@Id			BurgersCreate
////	@Accept		json
////	@Success	201
////	@Failure	400
////	@Failure	500
////	@Security	none
////	@Router		/missions/{m_id}/Burgers [post]
////	@Param		m_id	path	integer	true	"Mission ID"
////	@Param		json	body	request.Target	true	"Target data"
//func (h *burgersHandles) Create(ctx *gin.Context) {
//	uri, err := request.BindMissionUri(ctx)
//	if err != nil {
//		return
//	}
//	var req request.Target
//	if err := request.BindJSON(ctx, &req); err != nil {
//		return
//	}
//	if err := h.dao.Create(ctx, &model.Target{TmID: uri.MId, TName: req.TName}); err != nil {
//		resp.Abort500(ctx, err)
//		return
//	}
//	ctx.Status(http.StatusCreated)
//}
//
//// Read
////
////	@Summary	get Target
////	@Tags		Burgers
////	@Id			BurgersRead
////	@Produce	json
////	@Success	200	{object}	model.Target	"Target data"
////	@Failure	400
////	@Failure	404
////	@Failure	500
////	@Security	none
////	@Router		/target/{t_id} [get]
////	@Param		t_id	path	integer	true	"Target ID"
//func (h *burgersHandles) Read(ctx *gin.Context) {
//	uri, err := request.BindTargetUri(ctx)
//	if err != nil {
//		return
//	}
//	pr, err := h.dao.Read(ctx, uri.TId)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			resp.Abort404NotFound(ctx, err)
//			return
//		}
//		resp.Abort500(ctx, err)
//		return
//	}
//	resp.JSON(ctx, http.StatusOK, pr)
//}
//
//// Update
////
////	@Summary	update Target
////	@Tags		Burgers
////	@Id			BurgersUpdate
////	@Accept		json
////	@Success	200
////	@Failure	400
////	@Failure	500
////	@Security	none
////	@Router		/target/{t_id} [put]
////	@Param		t_id	path	integer			true	"Target ID"
////	@Param		json	body	request.Target	true	"Target data"
//func (h *burgersHandles) Update(ctx *gin.Context) {
//	uri, err := request.BindTargetUri(ctx)
//	if err != nil {
//		return
//	}
//	var req request.Target
//	if err := request.BindJSON(ctx, &req); err != nil {
//		return
//	}
//	err = h.dao.Update(ctx, &model.Target{
//		TID:      uri.TId,
//		TName:    req.TName,
//		TCountry: req.TCountry,
//		TNote:    req.TNote,
//	})
//	if err != nil {
//		resp.Abort500(ctx, err)
//	}
//}
//
//// Delete
////
////	@Summary	delete Target
////	@Tags		Burgers
////	@Id			BurgersDelete
////	@Success	204
////	@Failure	400
////	@Failure	500
////	@Security	none
////	@Router		/target/{t_id} [delete]
////	@Param		t_id	path	integer	true	"Target ID"
//func (h *burgersHandles) Delete(ctx *gin.Context) {
//	uri, err := request.BindTargetUri(ctx)
//	if err != nil {
//		return
//	}
//	if err := h.dao.Delete(ctx, uri.TId); err != nil {
//		resp.Abort500(ctx, err)
//		return
//	}
//	ctx.Status(http.StatusNoContent)
//}
