package request

import (
	"app/internal/handlers/resp"
	"github.com/gin-gonic/gin"
	"time"
)

type SpyCatUri struct {
	CId int64 `uri:"c_id"  binding:"required" example:"1"`
}

func BindSpyCatUri(ctx *gin.Context) (*SpyCatUri, error) {
	var uri SpyCatUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type SpyCat struct {
	CId                int64     `json:"c_id"`
	CName              string    `json:"c_name" binding:"required,gte=1,lte=256"`
	CDtCreated         time.Time `json:"c_dt_created" example:"2020-01-01 01:01:01"`
	CYearsOfExperience int64     `json:"c_years_of_experience"`
	CBreed             string    `json:"c_breed"`
	CSalary            int64     `json:"c_salary" example:"10"`
}

type Login struct {
	Usr string `json:"usr" binding:"required,gte=1,lte=256" example:"admin"`
	Pwd string `json:"pwd" binding:"required,gte=1,lte=256" example:"123456"`
}

type MissionUri struct {
	MId int64 `uri:"m_id"  binding:"required" example:"1"`
}

func BindMissionUri(ctx *gin.Context) (*MissionUri, error) {
	var uri MissionUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type Mission struct {
	MName string `json:"m_name" binding:"required,gte=1,lte=256" example:"mission 1"`
}

type AssignSpyCatUri struct {
	MId int64 `uri:"m_id"  binding:"required" example:"1"`
	CId int64 `uri:"c_id"  binding:"required" example:"1"`
}

func BindAssignSpyCatUri(ctx *gin.Context) (*AssignSpyCatUri, error) {
	var uri AssignSpyCatUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type TargetUri struct {
	TId int64 `uri:"t_id"  binding:"required" example:"1"`
}

func BindTargetUri(ctx *gin.Context) (*TargetUri, error) {
	var uri TargetUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type Target struct {
	TName    string `json:"t_name" binding:"required,gte=1,lte=256" example:"target 1"`
	TCountry string `json:"t_country" example:"USA"`
	TNote    string `json:"t_note"  example:"NOTE 1"`
}
