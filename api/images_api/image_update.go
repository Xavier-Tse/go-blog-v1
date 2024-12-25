package images_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"选择文件id"`
	Name string `json:"name" binding:"required" msg:"输入文件名称"`
}

func (ImagesApi) ImagesUpdate(ctx *gin.Context) {
	var cr ImageUpdateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	var imageModel models.BannerModel
	if err := global.DB.Take(&imageModel, cr.ID).Error; err != nil {
		res.FailWithMessage("文件不存在", ctx)
		return
	}

	if err := global.DB.Model(&imageModel).Update("name", cr.Name).Error; err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("图片名称修改成功", ctx)
	return
}
