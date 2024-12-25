package menu_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdate(ctx *gin.Context) {
	var cr MenuRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	id := ctx.Param("id")

	// 先把之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", ctx)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()

	// 如果选择了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", ctx)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", ctx)
		return
	}

	res.OkWithMessage("修改菜单成功", ctx)

}
