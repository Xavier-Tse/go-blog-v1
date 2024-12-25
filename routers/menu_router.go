package routers

import (
	"Backend/api"
)

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menu", app.MenuCreate)
	router.GET("menu", app.MenuList)
	router.GET("menu_names", app.MenuNameList)
	router.PUT("menu/:id", app.MenuUpdate)
	router.DELETE("menu", app.MenuRemove)
	router.GET("menu/:id", app.MenuDetail)
}
