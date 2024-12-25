package message_api

import (
	"Backend/models"
	"Backend/models/res"
	"Backend/service/common"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListAll(ctx *gin.Context) {
	var cr models.PageInfo
	if err := ctx.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})

	res.OkWithList(list, count, ctx)
}
