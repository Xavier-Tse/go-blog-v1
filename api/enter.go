package api

import (
	"Backend/api/advertise_api"
	"Backend/api/article_api"
	"Backend/api/comment_api"
	"Backend/api/digg_api"
	"Backend/api/images_api"
	"Backend/api/menu_api"
	"Backend/api/message_api"
	"Backend/api/settings_api"
	"Backend/api/tag_api"
	"Backend/api/user_api"
)

type ApiGroup struct {
	SettingsApi  settings_api.SettingsApi
	ImagesApi    images_api.ImagesApi
	AdvertiseApi advertise_api.AdvertApi
	MenuApi      menu_api.MenuApi
	UserApi      user_api.UserApi
	TagApi       tag_api.TagApi
	MessageApi   message_api.MessageApi
	ArticleApi   article_api.ArticleApi
	Digg         digg_api.DiggApi
	Comment      comment_api.CommentApi
}

var ApiGroupApp = new(ApiGroup)
