package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysRoleApisRouter struct{}

func (r *SysRoleApisRouter) InitSysRoleApisRouter(Router *gin.RouterGroup) {
	sysRoleApisApi := v1.WebApiGroupApp.Sys.SysRoleApisApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 保存
		router.POST("/role/api/save", sysRoleApisApi.SaveData)
		// 查询明细 /user/get/1/xxx
		router.POST("/role/api/list", sysRoleApisApi.SelectRoleApis)
		// 角色改变
		router.POST("/role/api/change", sysRoleApisApi.ChangeRoleIdMenus)
	}
}
