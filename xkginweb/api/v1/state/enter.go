package state

import "xkginweb/service"

// 继承聚合的思想---聚合共享
type WebApiGroup struct {
	UserStateApi
}

// 公共实例---服务共享
var (
	userStatService = service.ServiceGroupApp.UserStateServiceGroup.UserStateService
)
