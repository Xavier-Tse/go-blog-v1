package ctype

import (
	"github.com/bytedance/sonic"
)

type Role int
type Status int

const (
	PermissionAdmin        Role = 1 // 管理员
	PermissionUser         Role = 2 // 普通用户
	PermissionVisitor      Role = 3 // 游客
	PermissionDisabledUser Role = 4 // 封禁用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisabledUser:
		str = "封禁用户"
	default:
		str = "其他"
	}
	return str
}
