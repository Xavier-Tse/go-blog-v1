package tag_api

import (
	"Backend/models"
	"Backend/models/res"
	"Backend/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagList(ctx *gin.Context) {
	var cr models.PageInfo
	if err := ctx.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, ctx)
}
