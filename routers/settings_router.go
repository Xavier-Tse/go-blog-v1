package routers

import (
	"Backend/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdate)
}
