package sys

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/model/entity/sys"
)

type SysUserRolesApi struct{}

// 用户授权角色
func (api *SysUserRolesApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	//roleIds, _ := c.GetQuery("roleIds")
	//if roleIds == "" {
	//	// 如果参数注入失败或者出错就返回接口调用这。出错了.
	//	response.FailWithMessage("请选择授权的角色", c)
	//	return
	//}

	// 我们没有使用上面的 c.GetQuery 的方式接收前端传过来的参数的哦，下面就用 role 对象的方式来获取 UserId 和 RoleIds 参数
	type roles struct {
		UserId  uint   `json:"userId"`
		RoleIds string `json:"roleIds"`
	}
	role := roles{}
	err := c.ShouldBindJSON(&role)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage("授权角色参数注入失败", c)
		return
	}

	// 开始对授权的角色进行分割
	//roleIdsSplit := strings.Split(roleIds, ",")
	//开始对授权的角色进行分割
	roleIdsSplit := strings.Split(role.RoleIds, ",")
	// 准备一个用户角色中间表的切片对象
	var sysUserRoles []*sys.SysUserRoles
	// 开始遍历
	for _, roleId := range roleIdsSplit {
		parseUint, _ := strconv.ParseUint(roleId, 10, 64)
		sysUserRole := sys.SysUserRoles{}
		//sysUserRole.UserId = c.GetUint("UserId")
		sysUserRole.UserId = role.UserId
		sysUserRole.RoleId = uint(parseUint)
		sysUserRoles = append(sysUserRoles, &sysUserRole)
	}

	// 创建实例，保存帖子
	//err2 := sysUserRolesService.SaveSysUserRoles(c.GetUint("userId"), sysUserRoles)
	err2 := sysUserRolesService.SaveSysUserRoles(role.UserId, sysUserRoles)
	// 如果保存失败。就返回创建失败的提升
	if err2 != nil {
		response.FailWithMessage("授权失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("授权成功", c)
}

/*
*
查询用户授权
*/
func (api *SysUserRolesApi) SelectUserRoles(c *gin.Context) {
	roles, _ := sysUserRolesService.SelectUserRoles(c.GetUint("userId"))
	response.Ok(roles, c)
}
