package advertise_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (AdvertApi) AdvertUpdateView(ctx *gin.Context) {
	var cr AdvertRequest
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", ctx)
		return
	}

	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", ctx)
		return
	}

	res.OkWithMessage("修改广告成功", ctx)
}
