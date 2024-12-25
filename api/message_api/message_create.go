package message_api

import (
	"Backend/global"
	"Backend/models"
	"Backend/models/res"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID    uint   `json:"send_user_id" binding:"required"`    // 发送人id
	ReceiveUserID uint   `json:"receive_user_id" binding:"required"` // 接收人id
	Content       string `json:"content" binding:"required"`         // 消息内容
}

// MessageCreate 发布消息
func (MessageApi) MessageCreate(ctx *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	var sendUser, receiveUser models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", ctx)
		return
	}
	err = global.DB.Take(&receiveUser, cr.ReceiveUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", ctx)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.ReceiveUserID,
		RevUserNickName:  receiveUser.NickName,
		RevUserAvatar:    receiveUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", ctx)
		return
	}
	res.OkWithMessage("消息发送成功", ctx)
	return
}
