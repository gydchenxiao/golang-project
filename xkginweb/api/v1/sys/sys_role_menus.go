package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/sys"
)

type SysRoleMenusApi struct {
	global.BaseApi
}

// 接收参数(前端：await SaveRoleMenus({ "roleId": roleId, "menuIds": menuIds }))
type SysRoleMenusContext struct {
	RoleId  uint   `json:"roleId" validate:"required"`
	MenuIds string `json:"menuIds" validate:"required"`
}

// Messages 您可以自定义验证器错误消息
func (f SysRoleMenusContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f SysRoleMenusContext) Translates() map[string]string {
	return validate.MS{
		"RoleId":  "角色ID",
		"MenuIds": "菜单ID",
	}
}

// 角色授权菜单
func (api *SysRoleMenusApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysRoleMenusContext SysRoleMenusContext
	err := c.ShouldBindJSON(&sysRoleMenusContext)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithBindParams(c)
		return
	}

	//2: 开始验证
	validation := validate.Struct(sysRoleMenusContext)
	if !validation.Validate() {
		response.FailWithValidatorData(validation, c)
		return
	}

	// 把前端传过来的菜单ids进行分割
	menuIdsSplit := strings.Split(sysRoleMenusContext.MenuIds, ",")
	var sysRoleMenus []*sys.SysRoleMenus
	for _, menuId := range menuIdsSplit {
		parseUint, _ := strconv.ParseUint(menuId, 10, 64)
		sysRoleMenu := sys.SysRoleMenus{}
		sysRoleMenu.RoleId = sysRoleMenusContext.RoleId
		sysRoleMenu.MenuId = uint(parseUint)
		sysRoleMenus = append(sysRoleMenus, &sysRoleMenu)
	}

	// 创建实例，保存帖子
	err2 := sysRoleMenusService.SaveSysRoleMenus(sysRoleMenusContext.RoleId, sysRoleMenus)
	// 如果保存失败。就返回创建失败的提升
	if err2 != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

/**查询用户角色对应的菜单*/
func (api *SysRoleMenusApi) SelectRoleMenus(c *gin.Context) {
	roleIdStr := c.Query("roleId")
	roleId := api.StringToUnit(roleIdStr)
	if roleId == 0 {
		response.FailWithMessage("请选择一个角色操作", c)
		return
	}
	roles, _ := sysRoleMenusService.SelectRoleMenus(roleId)
	response.Ok(roles, c)
}
