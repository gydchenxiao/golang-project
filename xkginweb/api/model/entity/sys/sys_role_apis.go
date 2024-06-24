package sys

type SysRoleApis struct {
	RoleId uint `gorm:"comment:角色ID"`
	ApiId  uint `gorm:"comment:ApiID"`
}

func (s *SysRoleApis) TableName() string {
	return "sys_role_apis"
}
