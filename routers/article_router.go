package routers

import (
	"Backend/api"
	"Backend/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreate)
	router.GET("articles", app.ArticleList)
	router.GET("articles/detail", app.ArticleDetailByTitle)
	router.GET("articles/calendar", app.ArticleCalendar)
	router.GET("articles/tags", app.ArticleTagList)
	router.PUT("articles", middleware.JwtAuth(), app.ArticleUpdate)
	router.DELETE("articles", middleware.JwtAuth(), app.ArticleRemove)
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreate)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollectList)
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollectBatchRemove)
	router.GET("articles/text", app.FullTextSearch)
	router.GET("articles/:id", app.ArticleDetail)
}
