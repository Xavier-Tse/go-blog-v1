package settings_api

import (
	"Backend/global"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项配置信息
func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	var cr SettingsUri
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, ctx)
	case "email":
		res.OkWithData(global.Config.Email, ctx)
	case "qq":
		res.OkWithData(global.Config.QQ, ctx)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, ctx)
	case "jwt":
		res.OkWithData(global.Config.Jwt, ctx)
	default:
		res.FailWithMessage("没有对应配置信息", ctx)
	}
}
