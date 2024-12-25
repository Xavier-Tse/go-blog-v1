package user_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/ctype"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRole 用户权限变更
func (UserApi) UserUpdateRole(ctx *gin.Context) {
	var cr UserRole
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户id错误，用户不存在", ctx)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", ctx)
		return
	}
	res.OkWithMessage("修改权限成功", ctx)
}
