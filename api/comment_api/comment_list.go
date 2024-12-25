package comment_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/redis_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

func (CommentApi) CommentList(ctx *gin.Context) {
	var cr CommentListRequest
	err := ctx.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	res.OkWithData(filter.Select("c", rootCommentList), ctx)
	return
}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	// 遍历根评论，递归查根评论下的所有子评论
	diggInfo := redis_service.NewCommentDigg().GetInfo()
	for _, model := range RootCommentList {
		var subCommentList, newCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		for _, comment := range subCommentList {
			digg := diggInfo[fmt.Sprintf("%d", comment.ID)]
			comment.DiggCount += digg
			newCommentList = append(newCommentList, comment)
		}
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount += modelDigg
		model.SubComments = subCommentList
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}

// FindSubCommentCount 递归查评论下的子评论
func FindSubCommentCount(model models.CommentModel) (subCommentList []models.CommentModel) {
	findSubCommentList(model, &subCommentList)
	return subCommentList
}

func findSubCommentList(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}
