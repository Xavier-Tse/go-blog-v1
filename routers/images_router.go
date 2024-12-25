package routers

import (
	"Backend/api"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("image_names", imagesApi.ImageNameList)
	router.DELETE("images", imagesApi.ImageRemove)
	router.PUT("images", imagesApi.ImagesUpdate)
}
