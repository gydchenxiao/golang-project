package sys

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/sys"
)

type SysApisApi struct {
	global.BaseApi
}

// 拷贝
func (api *SysApisApi) CopyData(c *gin.Context) {
	// 1: 获取id数据 注意定义李媛媛的/:id
	id := c.Param("id")
	data, _ := sysApisService.CopyData(api.StringToUnit(id))
	response.Ok(data, c)
}

// 保存
func (api *SysApisApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysApis sys.SysApis
	err := c.ShouldBindJSON(&sysApis)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建实例，保存帖子
	err = sysApisService.SaveSysApis(&sysApis)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

// 状态修改
func (api *SysApisApi) UpdateStatus(c *gin.Context) {
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

	err = sysApisService.UpdateStatus(params.Id, params.Filed, params.Value)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 编辑修改
func (api *SysApisApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysApis sys.SysApis
	err := c.ShouldBindJSON(&sysApis)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结构体转化成map呢？
	m := structs.Map(sysApis)
	m["is_deleted"] = sysApis.IsDeleted
	err = sysApisService.UpdateSysApisMap(&sysApis, &m)
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
func (api *SysApisApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := sysApisService.DelSysApisById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 根据id查询信息
func (api *SysApisApi) GetById(c *gin.Context) {
	// 根据id查询方法
	id := c.Param("id")
	// 根据id查询方法
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	sysUser, err := sysApisService.GetSysApisByID(uint(parseUint))
	if err != nil {
		global.SugarLog.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}

	response.Ok(sysUser, c)
}

// 批量删除
func (api *SysApisApi) DeleteByIds(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	ids := c.Query("ids")
	idstrings := strings.Split(ids, ",")
	var sysApis []sys.SysApis
	for _, id := range idstrings {
		parseUint, _ := strconv.ParseUint(id, 10, 64)
		sysApi := sys.SysApis{}
		sysApi.ID = uint(parseUint)
		sysApis = append(sysApis, sysApi)
	}

	err := sysApisService.DeleteSysApissByIds(sysApis)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 分页查询信息
//
//	func (api *SysApisApi) LoadSysApisPage(c *gin.Context) {
//		// 创建一个分页对象
//		var pageInfo req.PageInfo
//		// 把前端json的参数传入给PageInfo
//		err := c.ShouldBindJSON(&pageInfo)
//		if err != nil {
//			response.FailWithMessage(err.Error(), c)
//			return
//		}
//
//		sysApisPage, total, err := sysApisService.LoadSysApisPage(pageInfo)
//		if err != nil {
//			response.FailWithMessage("获取失败"+err.Error(), c)
//			return
//		}
//
//		response.Ok(resp.PageResult{
//			List:     sysApisPage,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, c)
//	}
//
// 查询权限信息
func (api *SysApisApi) FindApisTree(c *gin.Context) {
	keyword := c.Query("keyword")
	sysApis, err := sysApisService.FinApiss(keyword)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(sysApisService.Tree(sysApis, 0), c)
}

// 查询父权限
func (api *SysApisApi) FindApisRoot(c *gin.Context) {
	sysMenus, err := sysApisService.FinApisRoot()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(sysMenus, c)
}
