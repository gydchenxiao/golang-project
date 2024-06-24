package bbs

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/bbs"
	"xkginweb/model/entity/bbs/request"
	resp "xkginweb/model/entity/comms/response"
	"xkginweb/utils"
)

// 定义api接口
type BbsCategoryApi struct{}

func (e *BbsCategoryApi) CreateBbsCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var saveReq request.BbsCategorySaveReq
	err := c.ShouldBindJSON(&saveReq)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	bbsCategory := new(bbs.BbsCategory)
	utils.CopyProperties(bbsCategory, saveReq)
	// 创建实例，保存帖子
	err = bbsCatgoryService.CreateBbsCategory(bbsCategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		global.Log.Error("创建失败")
		response.FailWithMessage("创建失败", c)
		return
	}

	global.Log.Info("创建成功")
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

func (e *BbsCategoryApi) UpdateBbsCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var bbsReq request.BbsCategorySaveReq
	err := c.ShouldBindJSON(&bbsReq)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 这里查询一次即可
	category, err := bbsCatgoryService.GetBbsCategory(bbsReq.ID)
	utils.CopyProperties(category, bbsReq)

	err = bbsCatgoryService.UpdateBbsCategory(category)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 更新状态
func (e *BbsCategoryApi) UpdateBbsCategoryStatus(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var statusReq request.StatusReq
	err := c.ShouldBindJSON(&statusReq)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = bbsCatgoryService.UpdateBbsCategoryStatus(&statusReq)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// GetBbsCategory
//
//	@Tags		GetBbsCategory
//	@Summary	根据ID查询帖子分类
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		bbs.GetBbsCategory													true	"客户ID"
//	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
//	@Router		/bbs/get?id=1 [get]
func (e *BbsCategoryApi) GetBbsCategory(c *gin.Context) {
	var bbsCategory bbs.BbsCategory
	// 绑定参数
	err := c.ShouldBindQuery(&bbsCategory)
	// 如果参数没有直接报错
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := bbsCatgoryService.GetBbsCategory(bbsCategory.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	response.Ok(data, c)
}

// FindBbsCategory
//
//	@Tags		FindBbsCategory
//	@Summary	查询所有有效的分类信息
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
//	@Router		/bbs/find  [get]
func (e *BbsCategoryApi) FindBbsCategory(c *gin.Context) {
	data, err := bbsCatgoryService.FindBbsCategory()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(data, c)
}

// http://localhost:8888/bbs/delete/:id
func (e *BbsCategoryApi) DeleteByBbsCategoryId(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := bbsCatgoryService.DeleteBbsCategoryById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// http://localhost:8888/bbs/deletes?ids=1,2,,3,4,5,6
func (e *BbsCategoryApi) DeleteByBbsCategoryIds(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	ids := c.Query("ids")
	// 开始执行
	idstrings := strings.Split(ids, ",")
	var bbscategories []bbs.BbsCategory
	for _, id := range idstrings {
		parseUint, _ := strconv.ParseUint(id, 10, 64)
		category := bbs.BbsCategory{}
		category.ID = uint(parseUint)
		bbscategories = append(bbscategories, category)
	}

	err := bbsCatgoryService.DeleteBbsCategoryByIds(bbscategories)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

func (e *BbsCategoryApi) GetBbsCategoryDetail(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 这个是用来获取?age=123
	//age := c.Query("age")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	data, err := bbsCatgoryService.GetBbsCategory(uint(parseUint))
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(data, c)
}

// get---params ----ShouldBindQuery
// post----data----ShouldBindJSON
// LoadBbsCategoryPage
func (e *BbsCategoryApi) LoadBbsCategoryPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo request.BbsCategoryPageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	BbsCategoryPage, total, err := bbsCatgoryService.LoadBbsCategoryPage(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.Ok(resp.PageResult{
		List:     BbsCategoryPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
