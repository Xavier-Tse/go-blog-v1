package comment_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/redis_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (CommentApi) CommentDigg(ctx *gin.Context) {
	var cr CommentIDRequest
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", ctx)
		return
	}

	redis_service.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))
	res.OkWithMessage("评论点赞成功", ctx)
}
