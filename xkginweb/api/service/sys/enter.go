package sys

// 继承聚合---方便统一管理和调用
type ServiceGroup struct {
	SysUserService
	SysMenusService
	SysApisService
	SysRolesService
	SysRoleMenusService
	SysRoleApisService
	SysUserRolesService
}
