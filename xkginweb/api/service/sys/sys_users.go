package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/comms/request"
	"xkginweb/model/entity/sys"
	"xkginweb/model/vo"
	"xkginweb/service/commons"
)

// 对用户表的数据层处理
type SysUserService struct {
	commons.BaseService[uint, sys.SysUser]
}

// 用于登录
func (service *SysUserService) GetUserByAccount(account string) (sysUser *sys.SysUser, err error) {
	// 根据account进行查询
	err = global.KSD_DB.Where("account = ?", account).First(&sysUser).Error
	if err != nil {
		return nil, err
	}
	return sysUser, nil
}

// 添加
func (service *SysUserService) SaveSysUser(sysUser *sys.SysUser) (err error) {
	err = global.KSD_DB.Create(sysUser).Error
	return err
}

// (系统用户启用和未启用、删除和未删除)状态修改
func (service *SysUserService) UpdateStatus(id uint, field string, fieldValue int) (err error) {
	var sysUser sys.SysUser
	err = global.KSD_DB.Unscoped().Model(&sysUser).
		Where("id = ?", id).
		Update(field, fieldValue).Error
	return err
}

// 更新
func (service *SysUserService) UpdateSysUser(sysUser *sys.SysUser) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysUser).Updates(sysUser).Error
	return err
}

// 按照Map的方式更新
func (service *SysUserService) UpdateSysUserMap(sysUser *sys.SysUser, mapFiled *map[string]any) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysUser).Updates(mapFiled).Error
	return err
}

// 删除
func (service *SysUserService) DelSysUserById(id uint) (err error) {
	var sysUser sys.SysUser
	err = global.KSD_DB.Where("id = ?", id).Delete(&sysUser).Error
	return err
}

// 批量删除
func (service *SysUserService) DeleteSysUsersByIds(sysUsers []sys.SysUser) (err error) {
	err = global.KSD_DB.Delete(&sysUsers).Error
	return err
}

// 根据id查询信息 (及得加上 Unscoped() 的哦)
func (service *SysUserService) GetSysUserByID(id uint) (sysUsers *sys.SysUser, err error) {
	err = global.KSD_DB.Unscoped().Where("id = ?", id).First(&sysUsers).Error
	return
}

//// 查询分页
//func (service *SysUserService) LoadSysUserPage(info request.PageInfo) (list interface{}, total int64, err error) {
//	// 获取分页的参数信息
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//
//	// 准备查询那个数据库表
//	db := global.KSD_DB.Model(&sys.SysUser{})
//
//	// 准备切片帖子数组
//	var sysUserList []sys.SysUser
//
//	// 加条件
//	if info.Keyword != "" {
//		db = db.Where("(username like ? or account like ? )", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
//	}
//
//	// 排序默时间降序降序
//	// 注意 .Unscoped() 的作用，就是去掉 IsDeleted soft_delete.DeletedAt 软删除(总是会存在sql语句中造成的影响)
//	db = db.Unscoped().Order("created_at desc")
//
//	// 查询中枢
//	err = db.Count(&total).Error
//	if err != nil {
//		return sysUserList, total, err
//	} else {
//		// 执行查询
//		err = db.Limit(limit).Offset(offset).Find(&sysUserList).Error
//	}
//
//	// 结果返回
//	return sysUserList, total, err
//}
//

// 加上了 RoleIds
// 查询分页
func (service *SysUserService) LoadSysUserPage(info request.PageInfo) (list interface{}, total int64, err error) {
	// 获取分页的参数信息
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 准备查询那个数据库表,这里为什么不使用Model呢，因为我要使用别名
	db := global.KSD_DB.Table("sys_users t1")

	// 加上一个 sys_user_roles 中间表的查询哦
	db.Select("t1.*,(SELECT GROUP_CONCAT(role_id)  FROM sys_user_roles WHERE user_id = t1.id) as RoleIds")
	// 准备切片帖子数组
	var sysUsersVos []vo.SysUsersVo
	// 加条件
	if info.Keyword != "" {
		db = db.Where("(t1.username like ? or t1.account like ? )", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}

	// 排序默时间降序降序
	db = db.Order("t1.created_at desc")

	// 查询中枢
	err = db.Unscoped().Count(&total).Error
	if err != nil {
		return sysUsersVos, total, err
	} else {
		// 执行查询
		err = db.Unscoped().Limit(limit).Offset(offset).Find(&sysUsersVos).Error
	}

	// 结果返回
	return sysUsersVos, total, err
}
