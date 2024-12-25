package middleware

import (
	"Backend/models/ctype"
	"Backend/models/res"
	"Backend/service/redis_service"
	"Backend/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", ctx)
			ctx.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", ctx)
			ctx.Abort()
			return
		}

		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token已失效", ctx)
			ctx.Abort()
			return
		}

		// 登录的用户
		ctx.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", ctx)
			ctx.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", ctx)
			ctx.Abort()
			return
		}

		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token已失效", ctx)
			ctx.Abort()
			return
		}

		// 登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", ctx)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
	}
}
