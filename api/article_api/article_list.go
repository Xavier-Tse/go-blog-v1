package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/es_service"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleList(ctx *gin.Context) {
	var cr ArticleSearchRequest
	if err := ctx.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, err := es_service.ComList(es_service.Option{
		PageInfo: cr.PageInfo,
		Field:    []string{"title", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", ctx)
		return
	}

	// 解决json-filter空值
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), ctx)
		return
	}
	res.OkWithList(data, int64(count), ctx)
}
