package service

import (
	"xkginweb/service/bbs"
	"xkginweb/service/state"
	"xkginweb/service/sys"
	"xkginweb/service/user"
	"xkginweb/service/video"
)

// 实例创建聚合
type ServicesGroup struct {
	SyserviceGroup        sys.ServiceGroup
	XkBbsServiceGroup     bbs.ServiceGroup
	XkVideoServiceGroup   video.ServiceGroup
	UserStateServiceGroup state.ServiceGroup
	UserServiceGroup      user.ServiceGroup
}

// 单例设计模式
var ServiceGroupApp = new(ServicesGroup)
