package tag_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

func (TagApi) TagCreate(ctx *gin.Context) {
	var cr TagRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	// 重复的判断
	var tag models.AdvertModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", ctx)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", ctx)
		return
	}

	res.OkWithMessage("添加标签成功", ctx)
}
