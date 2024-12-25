package images_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemove(ctx *gin.Context) {
	var cr models.RemoveRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", ctx)
		return
	}

	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), ctx)
}
