package routers

import (
	"Backend/api"
	"Backend/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("message", middleware.JwtAuth(), app.MessageCreate)
	router.GET("message_all", middleware.JwtAdmin(), app.MessageListAll)
	router.GET("message", middleware.JwtAuth(), app.MessageList)
}
