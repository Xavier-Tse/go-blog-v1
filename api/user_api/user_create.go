package user_api

import (
	"Backend/global"
	"Backend/models/ctype"
	"Backend/models/res"
	"Backend/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreate(ctx *gin.Context) {
	var cr UserCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	err := user_service.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", ctx.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), ctx)
	return
}
