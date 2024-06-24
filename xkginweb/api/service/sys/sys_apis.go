package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/sys"
	"xkginweb/service/commons"
)

// 对用户表的数据层处理
type SysApisService struct {
	commons.BaseService[uint, sys.SysApis]
}

// 数据复制
func (service *SysApisService) CopyData(id uint) (dbData *sys.SysApis, err error) {
	// 2: 查询id数据
	sysMenusData, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// 3: 开始复制
	sysMenusData.ID = 0
	// path 和 code 在数据库中设置成唯一索引了
	sysMenusData.Path += "/Index"
	sysMenusData.Code += "1"

	// 4: 保存入库
	data, err := service.Save(sysMenusData)

	return data, err
}

// 添加
func (service *SysApisService) SaveSysApis(sysApis *sys.SysApis) (err error) {
	err = global.KSD_DB.Create(sysApis).Error
	return err
}

// 修改
func (service *SysApisService) UpdateSysApis(sysApis *sys.SysApis) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysApis).Updates(sysApis).Error
	return err
}

// 按照map的方式更新
func (service *SysApisService) UpdateSysApisMap(sysApis *sys.SysApis, mapFileds *map[string]any) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysApis).Updates(mapFileds).Error
	return err
}

// 处理启用和未启用
func (service *SysApisService) UpdateStatus(id uint, field string, fieldValue int) (err error) {
	var sysApis sys.SysApis
	err = global.KSD_DB.Unscoped().Model(&sysApis).
		Where("id = ?", id).
		Update(field, fieldValue).Error
	return err
}

// 物理删除
func (service *SysApisService) DelSysApisById(id uint) (err error) {
	var sysApis sys.SysApis
	err = global.KSD_DB.Unscoped().Where("id = ?", id).Delete(&sysApis).Error
	return err
}

// 批量删除
func (service *SysApisService) DeleteSysApissByIds(sysApiss []sys.SysApis) (err error) {
	err = global.KSD_DB.Delete(&sysApiss).Error
	return err
}

// 根据id查询信息
func (service *SysApisService) GetSysApisByID(id uint) (sysApiss *sys.SysApis, err error) {
	err = global.KSD_DB.Unscoped().Omit("created_at", "updated_at").Where("id = ?", id).First(&sysApiss).Error
	return
}

// 查询分页
//
//	func (service *SysApisService) LoadSysApisPage(info request.PageInfo) (list interface{}, total int64, err error) {
//		// 获取分页的参数信息
//		limit := info.PageSize
//		offset := info.PageSize * (info.Page - 1)
//
//		// 准备查询那个数据库表,这里为什么不使用Model呢，因为我要使用别名
//		db := global.KSD_DB.Model(sys.SysApis{})
//
//		// 准备切片帖子数组
//		var sysApiss []sys.SysApis
//		// 加条件
//		if info.Keyword != "" {
//			db = db.Where("(title like ? or path like ?)", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
//		}
//
//		// 排序默时间降序降序
//		db = db.Order("created_at desc")
//
//		// 查询中枢
//		err = db.Unscoped().Count(&total).Error
//		if err != nil {
//			return sysApiss, total, err
//		} else {
//			// 执行查询
//			err = db.Unscoped().Limit(limit).Offset(offset).Find(&sysApiss).Error
//		}
//
//		// 结果返回
//		return sysApiss, total, err
//	}
//
// 查询权限信息
func (service *SysApisService) FinApiss(keyword string) (sysApis []*sys.SysApis, err error) {
	db := global.KSD_DB.Unscoped().Order("id asc")
	if len(keyword) > 0 {
		db.Where("title like ?", "%"+keyword+"%")
	}
	err = db.Find(&sysApis).Error
	return sysApis, err
}

/**
*   开始把数据进行编排--递归
*   Tree(all,0)
 */
func (service *SysApisService) Tree(allSysApis []*sys.SysApis, parentId uint) []*sys.SysApis {
	var nodes []*sys.SysApis
	for _, dbApis := range allSysApis {
		if dbApis.ParentId == parentId {
			childrensApis := service.Tree(allSysApis, dbApis.ID)
			if len(childrensApis) > 0 {
				dbApis.Children = append(dbApis.Children, childrensApis...)
			}
			nodes = append(nodes, dbApis)
		}
	}
	return nodes
}

/*
查询父级权限
*/
func (service *SysApisService) FinApisRoot() (sysApis []*sys.SysApis, err error) {
	err = global.KSD_DB.Where("parent_id = ? ", 0).Order("id asc").Find(&sysApis).Error
	return sysApis, err
}
