package bbs

import (
	"xkginweb/service"
)

// 继承聚合的思想---聚合共享
type WebApiGroup struct {
	XkBbsApi
	BbsCategoryApi
}

// 公共实例---服务共享
var (
	// 创建实例，保存帖子
	//xkBbsService      = new(bbs2.XkBbsService)
	//bbsCatgoryService = new(bbs2.BBSCategoryService)
	xkBbsService      = service.ServiceGroupApp.XkBbsServiceGroup.BbsService
	bbsCatgoryService = service.ServiceGroupApp.XkBbsServiceGroup.BbsCategoryService
)
