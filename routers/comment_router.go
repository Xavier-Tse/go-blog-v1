package routers

import (
	"Backend/api"
	"Backend/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.Comment
	router.POST("comments", middleware.JwtAuth(), app.CommentCreate)
	router.GET("comments/list", app.CommentList)
	router.GET("comments/:id", app.CommentDigg)
	router.DELETE("comments/:id", app.CommentRemove)
}
