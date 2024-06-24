package bbs

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type XkBbsRouter struct{}

func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {

	xkBbsApi := v1.WebApiGroupApp.Bbs.XkBbsApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("bbs")
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("save", xkBbsApi.CreateXkBbs)
		// 更新
		xkBbsCustomerRouterWithoutRecord.POST("update", xkBbsApi.UpdateXkBbs)
		// 更新
		xkBbsCustomerRouterWithoutRecord.DELETE("delete/:id", xkBbsApi.DeleteById)
		// 分页查询
		xkBbsCustomerRouterWithoutRecord.POST("page", xkBbsApi.LoadXkBbsPage)
		// 明细查询
		xkBbsCustomerRouterWithoutRecord.GET("get", xkBbsApi.GetXkBbs)
	}
}
