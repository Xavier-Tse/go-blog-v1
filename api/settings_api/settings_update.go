package settings_api

import (
	"Backend/config"
	"Backend/core"
	"Backend/global"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdate 修改某一项配置信息
func (SettingsApi) SettingsInfoUpdate(ctx *gin.Context) {
	var cr SettingsUri
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.SiteInfo = info

	case "email":
		var email config.Email
		err = ctx.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.Email = email

	case "qq":
		var qq config.QQ
		err = ctx.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.QQ = qq

	case "qiniu":
		var qiniu config.QiNiu
		err = ctx.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.QiNiu = qiniu

	case "jwt":
		var jwt config.Jwt
		err = ctx.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.Jwt = jwt

	default:
		res.FailWithMessage("没有对应配置信息", ctx)
		return
	}

	core.SetYaml()
	res.OkWith(ctx)
}
