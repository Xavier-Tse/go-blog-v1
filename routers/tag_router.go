package routers

import (
	"Backend/api"
)

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tags", app.TagCreate)
	router.GET("tags", app.TagList)
	router.PUT("tags/:id", app.TagUpdate)
	router.DELETE("tags", app.TagRemove)
}
