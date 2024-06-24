package video

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/model/entity/comms/request"
	resp "xkginweb/model/entity/comms/response"
	"xkginweb/model/entity/video"
)

// 定义api接口
type XkVideoCategoryApi struct{}

func (e *XkVideoCategoryApi) CreateXkVideoCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkcategory video.XkVideoCategory
	err := c.ShouldBindJSON(&xkcategory)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建实例，保存帖子
	err = xkcategoryService.CreateXkVideoCategory(&xkcategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

func (e *XkVideoCategoryApi) UpdateXkVideoCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkcategory video.XkVideoCategory
	err := c.ShouldBindJSON(&xkcategory)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = xkcategoryService.UpdateXkVideoCategory(&xkcategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

func (e *XkVideoCategoryApi) UpdateXkVideoCategoryStatus(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var statusReq request.StatusReq
	err := c.ShouldBindJSON(&statusReq)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = xkcategoryService.UpdateBbsCategoryStatus(&statusReq)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// GetXkVideoCategory
//
//	@Tags		GetXkVideoCategory
//	@Summary	根据ID查询帖子明细
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		video.GetXkVideoCategory													true	"客户ID"
//	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
//	@Router		/bbs/get?id=1 [get]
func (e *XkVideoCategoryApi) GetXkVideoCategory(c *gin.Context) {
	var xkcategory video.XkVideoCategory
	// 绑定参数
	err := c.ShouldBindQuery(&xkcategory)
	// 如果参数没有直接报错
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := xkcategoryService.GetXkVideoCategory(xkcategory.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	response.Ok(data, c)
}

// http://localhost:8888/bbs/delete/:id
func (e *XkVideoCategoryApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := xkcategoryService.DeleteXkVideoCategoryById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

func (e *XkVideoCategoryApi) GetXkVideoCategoryDetail(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 这个是用来获取?age=123
	//age := c.Query("age")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	data, err := xkcategoryService.GetXkVideoCategory(uint(parseUint))
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(data, c)
}

// LoadXkVideoCategoryPage
func (e *XkVideoCategoryApi) LoadXkVideoCategoryPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo request.PageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建实例，保存帖子
	xkcategoryPage, total, err := xkcategoryService.LoadXkVideoCategoryPage(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.Ok(resp.PageResult{
		List:     xkcategoryPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

// 查询视频分类信息
func (e *XkVideoCategoryApi) FindCategories(c *gin.Context) {
	catgories, err := xkcategoryService.FindCategories()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(xkcategoryService.Tree(catgories, 0), c)
}

// 查询所有的主分类信息
func (e *XkVideoCategoryApi) FindCategoryAll(c *gin.Context) {
	catgories, err := xkcategoryService.FindCategoryAll()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(catgories, c)
}

// http://localhost:8888/bbs/delete/:id
func (e *XkVideoCategoryApi) DeleteByVideoCategoryId(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := xkcategoryService.DeleteVideoCategoryById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// http://localhost:8888/bbs/deletes?ids=1,2,,3,4,5,6
func (e *XkVideoCategoryApi) DeleteByVideoCategoryIds(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	ids := c.Query("ids")
	// 开始执行
	idstrings := strings.Split(ids, ",")
	var videoCategoires []video.XkVideoCategory
	for _, id := range idstrings {
		parseUint, _ := strconv.ParseUint(id, 10, 64)
		category := video.XkVideoCategory{}
		category.ID = uint(parseUint)
		videoCategoires = append(videoCategoires, category)
	}

	err := xkcategoryService.DeleteVideoCategorysByIds(videoCategoires)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}
