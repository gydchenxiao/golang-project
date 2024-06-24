package sys

import "xkginweb/service"

type WebApiGroup struct {
	SysMenuApi
	SysUsersApi
	SysRolesApi
	SysApisApi
	SysUserRolesApi
	SysRoleMenusApi
	SysRoleApisApi
}

var (
	sysMenuService      = service.ServiceGroupApp.SyserviceGroup.SysMenusService
	sysUserService      = service.ServiceGroupApp.SyserviceGroup.SysUserService
	sysRolesService     = service.ServiceGroupApp.SyserviceGroup.SysRolesService
	sysApisService      = service.ServiceGroupApp.SyserviceGroup.SysApisService
	sysUserRolesService = service.ServiceGroupApp.SyserviceGroup.SysUserRolesService
	sysRoleApisService  = service.ServiceGroupApp.SyserviceGroup.SysRoleApisService
	sysRoleMenusService = service.ServiceGroupApp.SyserviceGroup.SysRoleMenusService
)
