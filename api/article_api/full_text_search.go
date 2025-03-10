package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"context"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) FullTextSearch(ctx *gin.Context) {
	var cr models.PageInfo
	_ = ctx.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value // 搜索到结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		sonic.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}
	res.OkWithList(fullTextList, count, ctx)
}
