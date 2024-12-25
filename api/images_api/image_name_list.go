package images_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImagesApi) ImageNameList(ctx *gin.Context) {
	var imageList []ImageResponse

	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, ctx)
}
