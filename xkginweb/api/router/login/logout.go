package login

import (
	"github.com/gin-gonic/gin"
	"xkginweb/api/v1/login"
)

// 登录路由
type LogoutRouter struct{}

func (router *LogoutRouter) InitLogoutRouter(Router *gin.RouterGroup) {
	logoutApi := login.LogOutApi{}
	// 用组定义--（推荐）
	loginRouter := Router.Group("/login")
	{
		loginRouter.POST("/logout", logoutApi.ToLogout)
	}
}
