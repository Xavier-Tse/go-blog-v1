package service

import (
	"Backend/service/image_service"
	"Backend/service/user_service"
)

type Group struct {
	ImageService image_service.ImageService
	userService  user_service.UserService
}

var GroupApp = new(Group)
