package routers

import (
	"Backend/api"
)

func (router RouterGroup) DiggRouter() {
	app := api.ApiGroupApp.Digg
	router.POST("digg/article", app.DiggArticle)
}
