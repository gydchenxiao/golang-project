package v1

import (
	"xkginweb/api/v1/bbs"
	"xkginweb/api/v1/code"
	"xkginweb/api/v1/course"
	"xkginweb/api/v1/login"
	"xkginweb/api/v1/state"
	"xkginweb/api/v1/sys"
	"xkginweb/api/v1/upload"
	"xkginweb/api/v1/video"
)

type WebApiGroup struct {
	Course course.WebApiGroup
	Video  video.WebApiGroup
	Code   code.WebApiGroup
	Sys    sys.WebApiGroup
	State  state.WebApiGroup
	Upload upload.WebApiGroup
	Bbs    bbs.WebApiGroup
	Login  login.WebApiGroup
}

var WebApiGroupApp = new(WebApiGroup)
