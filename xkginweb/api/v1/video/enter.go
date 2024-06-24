package video

import "xkginweb/service"

type WebApiGroup struct {
	XkVideoCategoryApi
	XkVideoApi
}

var (
	VALIDATOR_MAP        = map[string]string{"code": "701", "msg": "验证属性有误"}
	BINDING_PAMATERS_MAP = map[string]string{"code": "702", "msg": "参数有误"}
)
var (
	// 课程分类
	xkcategoryService = service.ServiceGroupApp.XkVideoServiceGroup.VideoCategoryService
	// 课程
	xkVideoService = service.ServiceGroupApp.XkVideoServiceGroup.VideoService
)
