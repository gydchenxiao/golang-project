package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysMenusRouter struct{}

func (r *SysMenusRouter) InitSysMenusRouter(Router *gin.RouterGroup) {
	sysMenuApi := v1.WebApiGroupApp.Sys.SysMenuApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 获取菜单列表
		router.POST("/menus/tree", sysMenuApi.FindMenus)
		// 获取父级权限
		router.POST("/menus/root", sysMenuApi.FindMenusRoot)
		// 保存
		router.POST("/menus/save", sysMenuApi.SaveData)
		// 修改
		router.POST("/menus/update", sysMenuApi.UpdateById)
		// 复制
		router.POST("/menus/copy/:id", sysMenuApi.CopyData)
		// 启用和未启用 （控制启用，发布，删除）
		router.POST("/menus/update/status", sysMenuApi.UpdateStatus)
		// 删除单个 :id 获取参数的时候id := c.Param("id")，传递的时候/sys/user/del/100
		router.POST("/menus/del/:id", sysMenuApi.DeleteById)
		// 查询明细 /user/get/1/xxx
		router.POST("/menus/get/:id", sysMenuApi.GetById)
	}
}
