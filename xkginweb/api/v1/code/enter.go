package code

import (
	"xkginweb/commons/jwtgo"
	"xkginweb/service"
)

// 继承聚合的思想---聚合共享
type WebApiGroup struct {
	CodeApi
}

// 公共实例---服务共享
var (
	sysUserService = service.ServiceGroupApp.SyserviceGroup.SysUserService
	jwtService     = jwtgo.JwtService{}
)
