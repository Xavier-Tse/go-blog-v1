package comment_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/es_service"
	"Backend/service/redis_service"
	"Backend/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

func (CommentApi) CommentCreate(ctx *gin.Context) {
	var cr CommentRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 文章是否存在
	_, err = es_service.CommDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMessage("文章不存在", ctx)
		return
	}

	// 判断是否是子评论
	if cr.ParentCommentID != nil {
		// 子评论
		// 给父评论数 +1
		// 父评论id
		var parentComment models.CommentModel
		// 找父评论
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMessage("父评论不存在", ctx)
			return
		}

		// 判断父评论的文章是否和当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("评论文章不一致", ctx)
			return
		}
		// 给父评论评论数+1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}

	// 添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	// 拿到文章数，新的文章评论数存缓存里
	// newCommentCount := article.CommentCount + 1
	// 给文章评论数 +1
	redis_service.NewCommentCount().Set(cr.ArticleID)
	res.OkWithMessage("文章评论成功", ctx)
	return
}
