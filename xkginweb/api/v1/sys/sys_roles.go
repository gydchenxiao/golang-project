package sys

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	req "xkginweb/model/entity/comms/request"
	resp "xkginweb/model/entity/comms/response"
	"xkginweb/model/entity/sys"
)

type SysRolesApi struct{}

// 查询用户角色
func (api *SysRolesApi) FindRoles(c *gin.Context) {
	roless, _ := sysRolesService.FindRoles()
	response.Ok(roless, c)
}

// 保存
func (api *SysRolesApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysRoles sys.SysRoles
	err := c.ShouldBindJSON(&sysRoles)
	if err != nil {
		fmt.Println(err)
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建实例，保存帖子
	err = sysRolesService.SaveSysRoles(&sysRoles)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

// 状态修改
func (api *SysRolesApi) UpdateStatus(c *gin.Context) {
	type Params struct {
		Id    uint   `json:"id"`
		Filed string `json:"field"`
		Value int    `json:"value"`
	}
	var params Params
	err := c.ShouldBindJSON(&params)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysRolesService.UpdateStatus(params.Id, params.Filed, params.Value)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 修改
func (api *SysRolesApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysRole sys.SysRoles
	err := c.ShouldBindJSON(&sysRole)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结构体转化成map呢？
	m := structs.Map(sysRole)
	m["is_deleted"] = sysRole.IsDeleted // global.GVA_MODEL `structs:"-"` 所以需要再加上的哦
	err = sysRolesService.UpdateSysRolesMap(&sysRole, &m)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("更新失败", c)
		return
	}

	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 根据id删除
func (api *SysRolesApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := sysRolesService.DelSysRolesById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 批量删除
func (api *SysRolesApi) DeleteByIds(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	ids := c.Query("ids")
	idstrings := strings.Split(ids, ",")
	var sysRoles []sys.SysRoles
	for _, id := range idstrings {
		parseUint, _ := strconv.ParseUint(id, 10, 64)
		sysRole := sys.SysRoles{}
		sysRole.ID = uint(parseUint)
		sysRoles = append(sysRoles, sysRole)
	}

	err := sysRolesService.DeleteSysRolessByIds(sysRoles)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 根据id查询信息
func (api *SysRolesApi) GetById(c *gin.Context) {
	// 根据id查询方法
	id := c.Param("id")
	// 根据id查询方法
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	sysUser, err := sysRolesService.GetSysRolesByID(uint(parseUint))
	if err != nil {
		global.SugarLog.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}

	response.Ok(sysUser, c)
}

// 分页查询信息
func (api *SysRolesApi) LoadSysRolesPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo req.PageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	sysRolesPage, total, err := sysRolesService.LoadSysRolesPage(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.Ok(resp.PageResult{
		List:     sysRolesPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
