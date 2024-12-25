package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/es_service"
	"Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

// ArticleCollCreate 用户收藏文章，或取消收藏
func (ArticleApi) ArticleCollCreate(ctx *gin.Context) {
	var cr models.ESIDRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	model, err := es_service.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", ctx)
		return
	}

	var collects models.UserCollectModel
	err = global.DB.Take(&collects, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	}
	// 取消收藏
	// 文章数 -1
	global.DB.Delete(&collects)

	// 更新文章收藏数
	err = es_service.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", ctx)
	} else {
		res.OkWithMessage("取消收藏成功", ctx)
	}
}
