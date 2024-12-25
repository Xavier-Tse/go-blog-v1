package routers

import (
	"Backend/api"
)

func (router RouterGroup) AdvertiseRouter() {
	app := api.ApiGroupApp.AdvertiseApi
	router.POST("advertise", app.AdvertCreateView)
	router.GET("advertise", app.AdvertListView)
	router.PUT("advertise", app.AdvertUpdateView)
	router.DELETE("advertise", app.AdvertRemoveView)
}
