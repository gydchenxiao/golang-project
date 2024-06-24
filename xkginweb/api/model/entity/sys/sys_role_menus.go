package sys

type SysRoleMenus struct {
	RoleId uint `gorm:"comment:角色ID"`
	MenuId uint `gorm:"comment:菜单ID"`
}

func (s *SysRoleMenus) TableName() string {
	return "sys_role_menus"
}
