package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/service/common"
	"Backend/utils/jwts"
	"context"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (ArticleApi) ArticleCollectList(ctx *gin.Context) {
	var cr models.PageInfo
	ctx.ShouldBindQuery(&cr)

	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var articleIDList []interface{}
	list, count, err := common.ComList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	var collectMap = map[string]string{}
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collectMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)
	var collList = make([]CollResponse, 0)

	// 传id列表，查es
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = sonic.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreatedAt:    collectMap[hit.Id],
		})
	}
	res.OkWithList(collList, count, ctx)
}
