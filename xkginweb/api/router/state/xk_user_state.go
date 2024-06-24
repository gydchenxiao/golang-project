package state

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type UserStateRouter struct{}

func (e *UserStateRouter) InitUserStateRouter(Router *gin.RouterGroup) {

	userStateApi := v1.WebApiGroupApp.State.UserStateApi
	xkBbsCustomerRouterWithoutRecord := Router.Group("state") //.Use(middleware.OperationRecord())
	{
		// 统计某年的用户注册量
		xkBbsCustomerRouterWithoutRecord.GET("user/reg", userStateApi.UserRegState)
		// 统计某年的用户注册量--明细信息
		xkBbsCustomerRouterWithoutRecord.POST("user/detail", userStateApi.FindUserRegStateDetail)
	}

}
