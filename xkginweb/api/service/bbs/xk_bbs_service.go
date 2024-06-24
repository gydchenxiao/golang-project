package bbs

import (
	"xkginweb/global"
	"xkginweb/model/entity/bbs"
	"xkginweb/model/entity/bbs/request"
)

// 定义bbs的service提供xkbbs的数据curd的操作

type BbsService struct{}

// @author: feige
// @function: CreateXkBbs
// @description: 创建文章
// @param: e bbs.XkBbs
// @return: err error
func (cbbs *BbsService) CreateXkBbs(xkBbs *bbs.XkBbs) (err error) {
	// 1： 获取数据的连接对象 如果执行成功err是nil，如果失败就把失败告诉
	err = global.KSD_DB.Create(xkBbs).Error
	return err
}

//@author: feige
//@function: UpdateXkBbs
//@description: 更新文章
//@param: e *model.ExaCustomer
//@return: err error

func (cbbs *BbsService) UpdateXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.KSD_DB.Save(xkBbs).Error
	//err = global.KSD_DB.Model(xkBbs).Updates(xkBbs).Error
	return err
}

// @author: feige
// @function: DeleteXkBbs
// @description: 删除帖子
// @param: e model.DeleteXkBbs
// @return: err error
func (cbbs *BbsService) DeleteXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.KSD_DB.Delete(&xkBbs).Error
	return err
}

// @author: feige
// @function: DeleteXkBbsById
// @description: 根据ID删除帖子
// @param: e model.DeleteXkBbsById
// @return: err error
func (cbbs *BbsService) DeleteXkBbsById(id uint) (err error) {
	var xkBbs bbs.XkBbs
	err = global.KSD_DB.Where("id = ?", id).Delete(&xkBbs).Error
	return err
}

// @author: feige
// @function: GetXkBbs
// @description: 根据ID获取帖子信息
// @param: id uint
// @return: xkBbs *bbs.XkBbs, err error
func (cbbs *BbsService) GetXkBbs(id uint) (xkBbs *bbs.XkBbs, err error) {
	err = global.KSD_DB.Where("id = ?", id).First(&xkBbs).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (cbbs *BbsService) LoadXkBbsPage(info request.BbsPageInfo) (list interface{}, total int64, err error) {
	// 获取分页的参数信息
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 准备查询那个数据库表
	db := global.KSD_DB.Model(&bbs.XkBbs{})

	// 准备切片帖子数组
	var XkBbsList []bbs.XkBbs

	if info.CategoryId != -1 {
		db = db.Where("category_id = ?", info.CategoryId)
	}

	if info.Status != -1 {
		db = db.Where("status = ?", info.Status)
	}

	// 加条件
	if info.Keyword != "" {
		db = db.Where("(title like ?  or category_name like ? or user_id like ? or username like ?)", "%"+info.Keyword+"%")
	}

	// 排序默时间降序降序
	db = db.Order("created_at desc")

	// 查询中枢
	err = db.Count(&total).Error
	if err != nil {
		return XkBbsList, total, err
	} else {
		// 执行查询
		err = db.Limit(limit).Offset(offset).Find(&XkBbsList).Error
	}

	// 结果返回
	return XkBbsList, total, err
}
