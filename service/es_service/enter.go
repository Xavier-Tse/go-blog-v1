package es_service

import (
	"Backend/models"
)

type Option struct {
	models.PageInfo
	Field []string
	Tag   string
}

func (o *Option) GetFrom() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
