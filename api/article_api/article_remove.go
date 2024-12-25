package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/es_service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemove(ctx *gin.Context) {
	var cr IDListRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	// TODO:文章删除了，但是有用户收藏
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		go es_service.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", ctx)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), ctx)
	return
}
