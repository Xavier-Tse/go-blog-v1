package user_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/utils/jwts"
	"Backend/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLogin(ctx *gin.Context) {
	var cr EmailLoginRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", ctx)
		return
	}

	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMessage("用户名或密码错误", ctx)
		return
	}

	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", ctx)
		return
	}

	res.OkWithData(token, ctx)
}
