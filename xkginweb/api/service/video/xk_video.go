package video

import (
	"xkginweb/global"
	"xkginweb/model/entity/video"
	"xkginweb/model/entity/video/request"
	"xkginweb/model/entity/video/response"
)

type VideoService struct {
}

// 保存
func (videoService *VideoService) SaveVideo(xkVideo *video.XkVideo) (err error) {
	err = global.KSD_DB.Model(xkVideo).Create(xkVideo).Error
	return err
}

// 更新
func (videoService *VideoService) UpdateVideo(xkVideo *video.XkVideo) (err error) {
	err = global.KSD_DB.Model(xkVideo).Updates(xkVideo).Error
	return err
}

// 根据id查询视频信息
func (videoService *VideoService) GetXkVideoById(id uint) (xkVideo *video.XkVideo, err error) {
	err = global.KSD_DB.Where("id = ?", id).Find(&xkVideo).Error
	return
}

// 查询所有的视频的数据并分页
func (videoService *VideoService) FindVideosPage(params *request.XkVideoReq) (resp response.XkVideoResp, err error) {
	var xkVideo video.XkVideo
	var xkVideos []video.XkVideo
	// Omit方法是把不需要的过滤掉
	db := global.KSD_DB.Model(xkVideo)
	// 状态控制
	if params.Status != -1 {
		db.Where("status = ?", params.Status)
	}

	// 分类查询---单个查询
	//if params.CategoryId != -1 && params.CategoryCid != -1 {
	//	db.Where("category_id = ? and category_cid = ?", params.CategoryId, params.CategoryCid)
	//}

	// 分类查询---多个分类查询
	if params.CategoryId != -1 {
		db.Where("FIND_IN_SET(?,category_pid) > 0", params.CategoryId)
	}

	if params.CategoryCid != -1 {
		db.Where("FIND_IN_SET(?,category_cid) > 0 ", params.CategoryCid)
	}

	// 根据时间时间分段查询
	if len(params.StartTime) > 0 && len(params.EndTime) > 0 {
		db.Where("create_time between ? and ?", params.StartTime, params.EndTime)
	}

	// 增加查询条件
	if len(params.Keyword) > 0 {
		db.Where("(title like ? or tags like ? or category_name like ? or category_cname like ?)", "%"+params.Keyword+"%", "%"+params.Keyword+"%", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
	}

	var total int64
	// 这里是count返回
	err = db.Count(&total).Error
	// 这里是分页查询处理 0 10 / 10 10 /20 10 /30 10 /40 10
	offset := (params.PageNum - 1) * params.PageSize
	err = db.Omit("content").Offset(offset).Limit(params.PageSize).Order("create_time desc").Find(&xkVideos).Error

	// 开始返回
	resp.Total = total
	resp.PageNum = params.PageNum
	resp.PageSize = params.PageSize
	resp.List = &xkVideos
	return resp, err
}
