package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/sys"
	"xkginweb/service/commons"
)

// 定义bbs的service提供BbsCategory的数据curd的操作

// 对用户表的数据层处理
type SysMenusService struct {
	commons.BaseService[uint, sys.SysMenus]
}

// 数据复制
func (service *SysMenusService) CopyData(id uint) (dbData *sys.SysMenus, err error) {
	// 2: 查询id数据
	sysMenusData, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// 3: 开始复制
	sysMenusData.ID = 0
	// path 和 code 在数据库中设置成唯一索引了
	sysMenusData.Path += "/Index"

	// 4: 保存入库
	data, err := service.Save(sysMenusData)

	return data, err
}

// 递归形成tree
func (service *SysMenusService) FindMenus(keyword string) (sysMenus []*sys.SysMenus, err error) {
	db := global.KSD_DB.Unscoped().Order("sort asc")
	if len(keyword) > 0 {
		db.Where("title like ?", "%"+keyword+"%")
	}
	err = db.Find(&sysMenus).Error
	return sysMenus, err
}

func (service *SysMenusService) Tree(allDbMenus []*sys.SysMenus, parentId uint) []*sys.SysMenus {
	var nodes []*sys.SysMenus //---------准备空教室
	// 开始遍历父类
	for _, dbMenu := range allDbMenus {
		if dbMenu.ParentId == parentId {
			dbMenu.Children = append(dbMenu.Children, service.Tree(allDbMenus, dbMenu.ID)...)
			nodes = append(nodes, dbMenu)
		}
	}
	return nodes
}

// 获取父级菜单(不加上 .Unscoped() 会自动加上 Where is_deleted = 0 的查找条件，是需要的 ---> 删除了的菜单就不用添加子菜单了)
func (service *SysMenusService) FindMenusRoot() (sysMenus []*sys.SysMenus, err error) {
	err = global.KSD_DB.Where("parent_id = ?", 0).Order("sort asc").Find(&sysMenus).Error
	return sysMenus, err
}

// 添加
func (service *SysMenusService) SaveSysMenus(sysMenus *sys.SysMenus) (err error) {
	err = global.KSD_DB.Create(sysMenus).Error
	return err
}

// 修改
func (service *SysMenusService) UpdateSysMenus(sysMenus *sys.SysMenus) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysMenus).Updates(sysMenus).Error
	return err
}

// 按照map的方式过呢更新
func (service *SysMenusService) UpdateSysMenusMap(sysMenus *sys.SysMenus, sysMenusMap *map[string]any) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysMenus).Updates(sysMenusMap).Error
	return err
}

// 处理启用和未启用
func (service *SysMenusService) UpdateStatus(id uint, field string, fieldValue int) (err error) {
	var sysMenus sys.SysMenus
	err = global.KSD_DB.Unscoped().Model(&sysMenus).
		Where("id = ?", id).
		Update(field, fieldValue).Error
	return err
}

// 删除
func (service *SysMenusService) DelSysMenusById(id uint) (err error) {
	var sysMenus sys.SysMenus
	err = global.KSD_DB.Where("id = ?", id).Delete(&sysMenus).Error
	return err
}

// 批量删除
func (service *SysMenusService) DeleteSysMenussByIds(sysMenuss []sys.SysMenus) (err error) {
	err = global.KSD_DB.Delete(&sysMenuss).Error
	return err
}

// 根据id查询信息
func (service *SysMenusService) GetSysMenusByID(id uint) (sysMenuss *sys.SysMenus, err error) {
	err = global.KSD_DB.Unscoped().Omit("created_at", "updated_at").Where("id = ?", id).First(&sysMenuss).Error
	return
}
