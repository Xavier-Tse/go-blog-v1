package images_api

import (
	"Backend/models"
	"Backend/models/res"
	"Backend/service/common"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageListView(ctx *gin.Context) {
	var cr models.PageInfo
	err := ctx.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	res.OkWithList(list, count, ctx)
}
