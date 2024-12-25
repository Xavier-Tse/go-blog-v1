package user_api

import (
	"Backend/global"
	"Backend/models/res"
	"Backend/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) UserLogout(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := ctx.Request.Header.Get("token")

	// 需要计算剩余过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Sub(now)

	fmt.Println(diff)
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", ctx)
		return
	}

	res.OkWithMessage("注销成功", ctx)
}
