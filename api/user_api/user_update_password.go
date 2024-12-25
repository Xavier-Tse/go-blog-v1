package user_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/utils/jwts"
	"Backend/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

// UserUpdatePassword 修改登录人的id
func (UserApi) UserUpdatePassword(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", ctx)
		return
	}

	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", ctx)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", ctx)
		return
	}
	res.OkWithMessage("密码修改成功", ctx)
	return
}
