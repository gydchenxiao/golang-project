package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/sys"
	"xkginweb/service/commons"
)

// 对用户表的数据层处理
type SysRoleApisService struct {
	commons.BaseService[uint, sys.SysRoleApis]
}

// 角色授予api
func (service *SysRoleApisService) SaveSysRoleApis(roleId uint, sysRolesApis []*sys.SysRoleApis) (err error) {
	tx := global.KSD_DB.Begin()
	// 删除用户对应的角色
	if err := tx.Where("role_id = ?", roleId).Delete(&sys.SysRoleApis{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 开始保存用户和角色的关系
	if err := tx.Create(sysRolesApis).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 查询角色授权的信息
func (service *SysRoleApisService) SelectRoleApis(roleId uint) (sysApiss []*sys.SysApis, err error) {
	err = global.KSD_DB.Select("t2.*").Table("sys_role_apis t1,sys_apis t2").
		Where("t1.api_id = t2.id AND t1.role_id = ? and t2.is_deleted = 0", roleId).Scan(&sysApiss).Error
	return sysApiss, err
}
