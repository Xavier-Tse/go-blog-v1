package article_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/ctype"
	"Backend/models/res"
	"Backend/service/es_service"
	"Backend/utils/jwts"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"请输入文章标题"`   // 文章标题
	Tags     ctype.Array `json:"tags"`                                     // 文章标签
	Abstract string      `json:"abstract"`                                 // 文章简介
	Content  string      `json:"content" binding:"required" msg:"请输入文章内容"` // 文章内容
	Category string      `json:"category"`                                 // 文章分类
	Source   string      `json:"source"`                                   // 文章来源
	Link     string      `json:"link"`                                     // 原文链接
	BannerID uint        `json:"banner_id"`                                // 文章封面id
}

func (ArticleApi) ArticleCreate(ctx *gin.Context) {
	var cr ArticleRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName

	// 校验content  xss

	// 处理content
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	// fmt.Println(doc.Text())
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstract == "" {
		// 汉字的截取不一样
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传banner_id,后台就随机去选择一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("没有banner数据", ctx)
			return
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// 查banner_id下的banner_url
	var bannerUrl string
	if err := global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error; err != nil {
		res.FailWithMessage("banner不存在", ctx)
		return
	}

	// 查用户头像
	var avatar string
	if err := global.DB.Model(models.UserModel{}).Where("id = ?", cr.BannerID).Select("avatar").Scan(&avatar).Error; err != nil {
		res.FailWithMessage("用户不存在", ctx)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Tags:         cr.Tags,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		Keyword:      cr.Title,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
	}

	if article.ISExistData() {
		res.FailWithMessage("文章已存在", ctx)
		return
	}

	if err := article.Create(); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	go es_service.SyncArticleByFullText(article.ID, article.Title, article.Content)
	res.OkWithMessage("文章发布成功", ctx)
}
