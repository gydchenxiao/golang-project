package sys

import (
	"xkginweb/global"
	sys2 "xkginweb/model/entity/sys"
	"xkginweb/model/vo"
	"xkginweb/service/commons"
)

// 对用户表的数据层处理
type SysUserRolesService struct {
	commons.BaseService[uint, sys2.SysUserRoles]
}

// 用户授权
func (service *SysUserRolesService) SaveSysUserRoles(userId uint, sysUserRoles []*sys2.SysUserRoles) (err error) {
	// 事务开启
	tx := global.KSD_DB.Begin()
	// 删除用户对应的角色-------------执行成功了，会立即提交吗？
	if err := tx.Where("user_id = ?", userId).Delete(&sys2.SysUserRoles{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 开始保存用户和角色的关系--------------执行成功了，会立即提交吗？
	if err := tx.Create(sysUserRoles).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 事务提交
	return tx.Commit().Error
}

// 查询用户授权的角色信息
func (service *SysUserRolesService) SelectUserRoles(userId uint) (sysRolesVos []*vo.SysRolesVo, err error) {
	err = global.KSD_DB.Select("t2.*").Table("sys_user_roles t1,sys_roles t2").
		Where("t1.user_id = ? and t1.role_id = t2.id", userId).Scan(&sysRolesVos).Error
	return sysRolesVos, err
}
