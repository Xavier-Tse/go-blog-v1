package user_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"Backend/plugins/email"
	"Backend/utils/jwts"
	"Backend/utils/pwd"
	"Backend/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/gomail.v2"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmail(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 用户绑定邮箱， 第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	session := sessions.Default(ctx)
	if cr.Code == nil {
		// 第一次，后台发验证码
		// 生成4位验证码， 将生成的验证码存入session
		code := random.Code(4)

		// 写入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", ctx)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是 "+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", ctx)
		return
	}

	// 第二次，用户输入邮箱，验证码，密码
	code := session.Get("valid_code")

	// 校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", ctx)
		return
	}

	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", ctx)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", ctx)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)

	// 第一次的邮箱，和第二次的邮箱也要做一致性校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", ctx)
		return
	}

	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", ctx)
	return
}