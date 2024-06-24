package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysRolesRouter struct{}

func (r *SysRolesRouter) InitSysRoleRouter(Router *gin.RouterGroup) {
	sysRolesApi := v1.WebApiGroupApp.Sys.SysRolesApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 查询全部角色
		router.POST("/role/list", sysRolesApi.FindRoles)
		// 保存
		router.POST("/role/save", sysRolesApi.SaveData)
		// 修改
		router.POST("/role/update", sysRolesApi.UpdateById)
		// 启用和未启用 （控制启用，发布，删除）
		router.POST("/role/update/status", sysRolesApi.UpdateStatus)
		// 删除单个 :id 获取参数的时候id := c.Param("id")，传递的时候/sys/user/del/100
		router.POST("/role/del/:id", sysRolesApi.DeleteById)
		// 删除多个  获取参数的时候ids := c.Query("ids")，传递的时候/sys/user/dels?ids=1,2,3,4
		router.POST("/role/dels", sysRolesApi.DeleteByIds)
		// 查询明细 /user/get/1/xxx
		router.POST("/role/get/:id", sysRolesApi.GetById)
		// 查询分页,搜索
		router.POST("/role/load", sysRolesApi.LoadSysRolesPage)
	}
}
