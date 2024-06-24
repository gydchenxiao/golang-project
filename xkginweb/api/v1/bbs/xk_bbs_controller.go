package bbs

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xkginweb/commons/response"
	"xkginweb/model/entity/bbs"
	"xkginweb/model/entity/bbs/request"
	resp "xkginweb/model/entity/comms/response"
)

// 定义api接口
type XkBbsApi struct{}

func (e *XkBbsApi) CreateXkBbs(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkBbs bbs.XkBbs
	err := c.ShouldBindJSON(&xkBbs)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建实例，保存帖子
	err = xkBbsService.CreateXkBbs(&xkBbs)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

func (e *XkBbsApi) UpdateXkBbs(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkBbs bbs.XkBbs
	err := c.ShouldBindJSON(&xkBbs)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = xkBbsService.UpdateXkBbs(&xkBbs)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// GetXkBbs
//
//	@Tags		GetXkBbs
//	@Summary	根据ID查询帖子明细
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		bbs.GetXkBbs													true	"客户ID"
//	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
//	@Router		/bbs/get?id=1 [get]
func (e *XkBbsApi) GetXkBbs(c *gin.Context) {
	var xkBbs bbs.XkBbs
	// 绑定参数
	err := c.ShouldBindQuery(&xkBbs)
	// 如果参数没有直接报错
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := xkBbsService.GetXkBbs(xkBbs.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	response.Ok(data, c)
}

// http://localhost:8888/bbs/delete/:id
func (e *XkBbsApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := xkBbsService.DeleteXkBbsById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

func (e *XkBbsApi) GetXkBbsDetail(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 这个是用来获取?age=123
	//age := c.Query("age")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	data, err := xkBbsService.GetXkBbs(uint(parseUint))
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(data, c)
}

// LoadXkBbsPage
func (e *XkBbsApi) LoadXkBbsPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo request.BbsPageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	xkBbsPage, total, err := xkBbsService.LoadXkBbsPage(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.Ok(resp.PageResult{
		List:     xkBbsPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
