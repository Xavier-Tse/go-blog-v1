package user_api

import (
	"Backend/models"
	"Backend/models/ctype"
	"Backend/models/res"
	"Backend/service/common"
	"Backend/utils/desens"
	"Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserList(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    true,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
			user.Password = ""
			user.IP = ""
			user.Addr = ""
			user.Tel = desens.DesensitizationTel(user.Tel)
			user.Email = desens.DesensitizationEmail(user.Email)
		}
		// 脱敏
		users = append(users, user)
	}

	res.OkWithList(users, count, ctx)
}
