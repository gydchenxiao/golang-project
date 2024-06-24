package router

import (
	"xkginweb/router/bbs"
	"xkginweb/router/course"
	"xkginweb/router/login"
	"xkginweb/router/state"
	"xkginweb/router/sys"
	"xkginweb/router/video"
)

type WebRouterGroup struct {
	Course  course.WebRouterGroup
	Video   video.WebRouterGroup
	SysMenu sys.WebRouterGroup
	State   state.WebRouterGroup
	BBs     bbs.WebRouterGroup
	Login   login.WebRouterGroup
}

var RouterWebGroupApp = new(WebRouterGroup)
