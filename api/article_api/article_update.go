package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/ctype"
	"Backend/models/res"
	"Backend/service/es_service"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`     // 文章标题
	Abstract string   `json:"abstract"`  // 文章简介
	Content  string   `json:"content"`   // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
	ID       string   `json:"id"`
}

func (ArticleApi) ArticleUpdate(ctx *gin.Context) {
	var cr ArticleUpdateRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, ctx)
		return
	}
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			res.FailWithMessage("banner不存在", ctx)
			return
		}
	}

	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}
	maps := structs.Map(&article)
	var DataMap = map[string]any{}

	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	err = article.GetDataByID(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("文章不存在", ctx)
		return
	}

	err = es_service.ArticleUpdate(cr.ID, DataMap)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("文章更新失败", ctx)
		return
	}

	// 更新成功，同步数据到全文搜索
	newArticle, _ := es_service.CommDetail(cr.ID)
	if article.Content != newArticle.Content || article.Title != newArticle.Title {
		es_service.DeleteFullTextByArticleID(cr.ID)
		es_service.SyncArticleByFullText(cr.ID, article.Title, newArticle.Content)
	}
	res.OkWithMessage("更新成功", ctx)
}
