package sys

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/sys"
)

type SysMenuApi struct {
	global.BaseApi
}

// 拷贝
func (api *SysMenuApi) CopyData(c *gin.Context) {
	// 1: 获取id数据 注意定义李媛媛的/:id
	id := c.Param("id")
	data, _ := sysMenuService.CopyData(api.StringToUnit(id))
	response.Ok(data, c)
}

// 查询菜单列表
func (api *SysMenuApi) FindMenus(c *gin.Context) {
	keyword := c.Query("keyword") // 获取 keyword 参数
	sysMenus, err := sysMenuService.FindMenus(keyword)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(sysMenuService.Tree(sysMenus, 0), c)
}

// 获取父级菜单
func (api *SysMenuApi) FindMenusRoot(c *gin.Context) {
	sysMenus, err := sysMenuService.FindMenusRoot()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(sysMenuService.Tree(sysMenus, 0), c)
}

// 保存
func (api *SysMenuApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysMenus sys.SysMenus
	err := c.ShouldBindJSON(&sysMenus)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建实例，保存帖子
	err = sysMenuService.SaveSysMenus(&sysMenus)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

// 状态修改
func (api *SysMenuApi) UpdateStatus(c *gin.Context) {
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

	err = sysMenuService.UpdateStatus(params.Id, params.Filed, params.Value)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 编辑修改
func (api *SysMenuApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysMenus sys.SysMenus
	err := c.ShouldBindJSON(&sysMenus)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结构体转化成map呢？
	m := structs.Map(sysMenus)
	m["is_deleted"] = sysMenus.IsDeleted //
	err = sysMenuService.UpdateSysMenusMap(&sysMenus, &m)
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
func (api *SysMenuApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	flag, _ := sysMenuService.UnDeleteByID(api.StringToUnit(id)) // 继承下来的方法（直接删除，不是软删除的哦）
	if flag == false {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 根据id查询信息
func (api *SysMenuApi) GetById(c *gin.Context) {
	// 根据id查询方法
	id := c.Param("id")
	// 根据id查询方法
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	sysUser, err := sysMenuService.GetSysMenusByID(uint(parseUint))
	if err != nil {
		global.SugarLog.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}

	response.Ok(sysUser, c)
}
