package flag

import (
	"Backend/global"
	"Backend/models"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FeedBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
