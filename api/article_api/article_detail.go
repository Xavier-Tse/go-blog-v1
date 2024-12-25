package article_api

import (
	"Backend/models/res"
	"Backend/service/es_service"
	"Backend/service/redis_service"
	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetail(ctx *gin.Context) {
	var cr ESIDRequest
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	redis_service.NewArticleLook().Set(cr.ID)
	model, err := es_service.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	res.OkWithData(model, ctx)
}

func (ArticleApi) ArticleDetailByTitle(ctx *gin.Context) {
	var cr ArticleDetailRequest
	err := ctx.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}
	model, err := es_service.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	res.OkWithData(model, ctx)
}
