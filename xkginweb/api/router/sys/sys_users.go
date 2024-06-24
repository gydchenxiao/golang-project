package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysUsersRouter struct{}

func (r *SysUsersRouter) InitSysUsersRouter(Router *gin.RouterGroup) {
	sysUsersApi := v1.WebApiGroupApp.Sys.SysUsersApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 保存
		router.POST("/user/save", sysUsersApi.SaveData)
		// 修改用户密码
		router.POST("/user/updatePwd", sysUsersApi.UpdatePwd)
		// 修改自己的密码
		router.POST("/user/update/self/pwd", sysUsersApi.UpdatePwdSelf)
		// 修改
		router.POST("/user/update", sysUsersApi.UpdateById)
		// (系统用户启用和未启用、删除和未删除)状态修改
		router.POST("/user/update/status", sysUsersApi.UpdateStatus)
		// 删除单个 :id 获取参数的时候id := c.Param("id")，传递的时候/sys/user/del/100
		router.POST("/user/del/:id", sysUsersApi.DeleteById)
		// 删除多个  获取参数的时候ids := c.Query("ids")，传递的时候/sys/user/dels?ids=1,2,3,4
		router.POST("/user/dels", sysUsersApi.DeleteByIds)
		// 查询明细 /user/get/1/xxx
		router.POST("/user/get/:id", sysUsersApi.GetById)
		// 查询分页,搜索
		router.POST("/user/load", sysUsersApi.LoadSysUserPage)
	}
}
