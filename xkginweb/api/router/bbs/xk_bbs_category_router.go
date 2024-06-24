package bbs

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type BBSCategoryRouter struct{}

func (e *BBSCategoryRouter) InitBBSCategoryRouter(Router *gin.RouterGroup) {

	bbsCategoryApi := v1.WebApiGroupApp.Bbs.BbsCategoryApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("bbscategory")
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("save", bbsCategoryApi.CreateBbsCategory)
		// 更新
		xkBbsCustomerRouterWithoutRecord.POST("update", bbsCategoryApi.UpdateBbsCategory)
		// 更新状态
		xkBbsCustomerRouterWithoutRecord.POST("update/status", bbsCategoryApi.UpdateBbsCategoryStatus)
		// 删除
		xkBbsCustomerRouterWithoutRecord.DELETE("delete/:id", bbsCategoryApi.DeleteByBbsCategoryId)
		// 批量删除
		xkBbsCustomerRouterWithoutRecord.DELETE("deletes", bbsCategoryApi.DeleteByBbsCategoryIds)
		// 分页查询
		xkBbsCustomerRouterWithoutRecord.POST("page", bbsCategoryApi.LoadBbsCategoryPage)
		// 明细查询
		xkBbsCustomerRouterWithoutRecord.GET("get", bbsCategoryApi.GetBbsCategory)
		// 明细查询
		xkBbsCustomerRouterWithoutRecord.GET("find", bbsCategoryApi.FindBbsCategory)
	}
}
