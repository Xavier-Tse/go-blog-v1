package models

import (
	"Backend/global"
	"Backend/models/ctype"
	"os"
)

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片类型，本地/七牛
}

func (b *BannerModel) AfterDelete() (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除数据库和本地文件
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
