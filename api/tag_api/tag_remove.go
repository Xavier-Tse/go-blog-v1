package tag_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemove(ctx *gin.Context) {
	var cr models.RemoveRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", ctx)
		return
	}
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), ctx)

}
