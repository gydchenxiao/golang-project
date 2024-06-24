package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"strconv"
	"xkginweb/commons/response"
	"xkginweb/model/entity/video"
	"xkginweb/model/entity/video/request"
)

// 定义api接口
type XkVideoApi struct{}

// 保存视频
func (e XkVideoApi) SaveVideo(c *gin.Context) {
	var xkVideo video.XkVideo
	err2 := c.ShouldBindJSON(&xkVideo)
	if err2 != nil {
		// 如果保存成功，就返回创建创建成功
		response.FailWithMessage("保存视频失败", c)
		return
	}

	//校验
	v := validate.Struct(xkVideo)
	// 如果验证失败
	if !v.Validate() {
		//all := v.Errors.All()
		response.FailWithMessage(VALIDATOR_MAP["msg"], c)
		return
	}

	err := xkVideoService.SaveVideo(&xkVideo)
	if err != nil {
		// 如果保存成功，就返回创建创建成功
		response.FailWithMessage("保存失败失败", c)
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok(xkVideo, c)
}

// 更新视频
func (e XkVideoApi) UpdateVideo(c *gin.Context) {

	var xkVideo video.XkVideo
	err2 := c.ShouldBindJSON(&xkVideo)
	if err2 != nil {
		// 如果保存成功，就返回创建创建成功
		response.FailWithMessage(BINDING_PAMATERS_MAP["msg"], c)
		return
	}

	err := xkVideoService.UpdateVideo(&xkVideo)
	if err != nil {
		// 如果保存成功，就返回创建创建成功
		response.FailWithMessage("保存失败失败", c)
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok(xkVideo, c)
}

// 查询分页
func (e *XkVideoApi) FindVideosPage(c *gin.Context) {
	params := new(request.XkVideoReq)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		response.FailWithMessage("绑定参数失败", c)
		return
	}

	videos, _ := xkVideoService.FindVideosPage(params)
	// 如果保存成功，就返回创建创建成功
	response.Ok(videos, c)
}

// 根据ID查询分类信息
func (e *XkVideoApi) GetVideosById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	videos, err := xkVideoService.GetXkVideoById(uint(parseUint))
	if err != nil {
		fmt.Println(err)
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok(videos, c)
}
