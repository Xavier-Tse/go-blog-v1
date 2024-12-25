package routers

import (
	"Backend/api"
	"Backend/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLogin)
	router.GET("users", middleware.JwtAuth(), app.UserList)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemove)
	router.POST("users", middleware.JwtAdmin(), app.UserCreate)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRole)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), app.UserLogout)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmail)
}
