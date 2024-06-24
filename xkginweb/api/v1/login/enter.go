package login

import (
	"xkginweb/service"
)

// 继承聚合的思想---聚合共享
type WebApiGroup struct {
	LoginApi
	LogOutApi
}

// 公共实例---服务共享
var (
	sysUserService      = service.ServiceGroupApp.SyserviceGroup.SysUserService
	sysUserRoleService  = service.ServiceGroupApp.SyserviceGroup.SysUserRolesService
	sysRoleMenusService = service.ServiceGroupApp.SyserviceGroup.SysRoleMenusService
	sysRoleApisService  = service.ServiceGroupApp.SyserviceGroup.SysRoleApisService
	sysMenuService      = service.ServiceGroupApp.SyserviceGroup.SysMenusService
	//jwtService     = jwtgo.JwtService{}
)
