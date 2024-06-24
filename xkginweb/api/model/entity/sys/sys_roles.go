package sys

import (
	"xkginweb/global"
)

// structs 属性的设置是为了：gorm框架更新多列中，更新0值失败的问题（因为修改了 is_deleted 软删除字段）
// 注意：structs 后跟的是数据库表中的字段名，并且作为 Map 结构的 key 值
type SysRoles struct {
	global.GVA_MODEL `structs:"-"`
	RoleName         string `json:"roleName" gorm:"comment:角色名" structs:"role_name"`  // 角色名
	RoleCode         string `json:"roleCode" gorm:"comment:角色代号" structs:"role_code"` // 角色代号
}

func (s *SysRoles) TableName() string {
	return "sys_roles"
}
