package tag_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdate(ctx *gin.Context) {
	var cr TagRequest
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", ctx)
		return
	}

	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", ctx)
		return
	}

	res.OkWithMessage("修改标签成功", ctx)
}
