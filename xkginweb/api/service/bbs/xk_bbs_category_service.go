package bbs

import (
	"xkginweb/global"
	"xkginweb/model/entity/bbs"
	"xkginweb/model/entity/bbs/request"
)

// 定义bbs的service提供BbsCategory的数据curd的操作
type BbsCategoryService struct{}

// 创建文字分类
func (cbbs *BbsCategoryService) CreateBbsCategory(bbsCategory *bbs.BbsCategory) (err error) {
	// 1： 获取数据的连接对象 如果执行成功err是nil，如果失败就把失败告诉
	err = global.KSD_DB.Create(bbsCategory).Error
	return err
}

// 更新文字分类
func (cbbs *BbsCategoryService) UpdateBbsCategory(bbsCategory *bbs.BbsCategory) (err error) {
	err = global.KSD_DB.Model(bbsCategory).Save(bbsCategory).Error
	return err
}

// 修改状态
func (cbbs *BbsCategoryService) UpdateBbsCategoryStatus(statusReq *request.StatusReq) (err error) {
	err = global.KSD_DB.Model(new(*bbs.BbsCategory)).Where("id=?", statusReq.ID).Update(statusReq.Field, statusReq.Value).Error
	return err
}

// 删除
func (cbbs *BbsCategoryService) DeleteBbsCategory(bbsCategory *bbs.BbsCategory) (err error) {
	err = global.KSD_DB.Delete(&bbsCategory).Error
	return err
}

// 删除
func (cbbs *BbsCategoryService) DeleteBbsCategoryById(id uint) (err error) {
	var BbsCategory bbs.BbsCategory
	err = global.KSD_DB.Where("id = ?", id).Delete(&BbsCategory).Error
	return err
}

// 批量删除
func (cbbs *BbsCategoryService) DeleteBbsCategoryByIds(bbsCategorys []bbs.BbsCategory) (err error) {
	err = global.KSD_DB.Delete(&bbsCategorys).Error
	return err
}

// 根据ID获取帖子信息
func (cbbs *BbsCategoryService) GetBbsCategory(id uint) (bbsCategory *bbs.BbsCategory, err error) {
	err = global.KSD_DB.Where("id = ?", id).First(&bbsCategory).Error
	return
}

// 根据ID获取帖子信息
func (cbbs *BbsCategoryService) FindBbsCategory() (bbsCategorys []bbs.BbsCategory, err error) {
	err = global.KSD_DB.Where("status = 1 and is_delete = 0").Find(&bbsCategorys).Error
	return
}

// 分页获取客户列表
func (cbbs *BbsCategoryService) LoadBbsCategoryPage(info request.BbsCategoryPageInfo) (list interface{}, total int64, err error) {
	// 获取分页的参数信息
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 准备查询那个数据库表
	db := global.KSD_DB.Model(&bbs.BbsCategory{})

	// 准备切片帖子数组
	var BbsCategoryList []bbs.BbsCategory

	// 加条件
	if info.Keyword != "" {
		db = db.Where("title like ?", "%"+info.Keyword+"%")
	}

	if info.Status != -1 {
		db = db.Where("status = ?", info.Status)
	}
	// 排序默时间降序降序
	db = db.Order("sorted asc")

	// 查询中枢
	err = db.Count(&total).Error
	if err != nil {
		return BbsCategoryList, total, err
	} else {
		// 执行查询
		err = db.Limit(limit).Offset(offset).Find(&BbsCategoryList).Error
	}

	// 结果返回
	return BbsCategoryList, total, err
}
