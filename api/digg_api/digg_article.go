package digg_api

import (
	"Backend/models"
	"Backend/models/res"
	"Backend/service/redis_service"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticle(ctx *gin.Context) {
	var cr models.ESIDRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	// 对长度校验
	// 查es
	redis_service.NewDigg().Set(cr.ID)
	res.OkWithMessage("文章点赞成功", ctx)
}
