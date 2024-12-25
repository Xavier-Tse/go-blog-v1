package common

import (
	"Backend/global"
	"Backend/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{
			// Logger: logger.Default.LogMode(logger.Info),
			Logger: global.MysqlLog,
		})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}

	// 使用相同的查询条件
	query := DB.Where(model)

	// 获取记录总数
	err = query.Model(&model).Count(&count).Error
	query = DB.Where(model)

	// 设置分页
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	// 执行查询
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
