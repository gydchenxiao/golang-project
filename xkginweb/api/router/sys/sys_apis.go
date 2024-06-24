package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysApisRouter struct{}

func (r *SysApisRouter) InitSysApisRouter(Router *gin.RouterGroup) {
	sysApisApi := v1.WebApiGroupApp.Sys.SysApisApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 获取菜单列表
		router.POST("/apis/tree", sysApisApi.FindApisTree)
		// 获取父级权限
		router.POST("/apis/root", sysApisApi.FindApisRoot)
		// 保存
		router.POST("/apis/save", sysApisApi.SaveData)
		// 复制数据
		router.POST("/apis/copy/:id", sysApisApi.CopyData)
		// 修改
		router.POST("/apis/update", sysApisApi.UpdateById)
		// 启用和未启用 （控制启用，发布，删除）
		router.POST("/apis/update/status", sysApisApi.UpdateStatus)
		// 删除单个 :id 获取参数的时候id := c.Param("id")，传递的时候/sys/user/del/100
		router.POST("/apis/del/:id", sysApisApi.DeleteById)
		// 删除多个  获取参数的时候ids := c.Query("ids")，传递的时候/sys/user/dels?ids=1,2,3,4
		router.POST("/apis/dels", sysApisApi.DeleteByIds)
		// 查询明细 /user/get/1/xxx
		router.POST("/apis/get/:id", sysApisApi.GetById)
	}
}
