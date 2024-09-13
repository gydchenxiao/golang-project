package service

import (
	"go-gorm-plus/global"
	"gorm.io/gorm"
)

// 父类
type BaseServiceImpl[D any, T any] struct{}

// 明细
func (service *BaseServiceImpl[D, T]) GetByID(id D) (data T, err error) {
	err = global.KSD_DB.Where("id = ?", id).First(&data).Error
	return
}

// 无逻辑条件
func (service *BaseServiceImpl[D, T]) UnGetByID(id D) (data T, err error) {
	err = global.KSD_DB.Unscoped().Where("id = ?", id).First(&data).Error
	return
}

// 删除
func (service *BaseServiceImpl[D, T]) DeleteByID(id D) (bool, int64) {
	var data T
	rows := global.KSD_DB.Where("id = ?", id).Delete(&data).RowsAffected
	return rows > 0, rows
}

// 忽略逻辑删除
func (service *BaseServiceImpl[D, T]) UnDeleteByID(id D) (bool, int64) {
	var data T
	rows := global.KSD_DB.Unscoped().Where("id = ?", id).Delete(&data).RowsAffected
	return rows > 0, rows
}

// 保存和更新
func (service *BaseServiceImpl[D, T]) Save(data T) (dbData *T, err error) {
	err = global.KSD_DB.Create(&data).Error
	return &data, err
}

// 批量保存
func (service *BaseServiceImpl[D, T]) SaveBatch(datas []T) (bool, int64) {
	affected := global.KSD_DB.Create(&datas).RowsAffected
	return affected > 0, affected
}

// 更新忽略物理条件
func (service *BaseServiceImpl[D, T]) UnUpdateByID(data T) (dbData *T, err error) {
	err = global.KSD_DB.Unscoped().Model(dbData).Updates(&data).Error
	return &data, err
}

// 更新逻辑更新
func (service *BaseServiceImpl[D, T]) UpdateByID(data T) (dbData *T, err error) {
	err = global.KSD_DB.Model(dbData).Updates(&data).Error
	return &data, err
}

// 更新逻辑更新
//func (service *BaseServiceImpl[D, T]) UpdateMap(data T) (dbData *T, err error) {
//	err = global.KSD_DB.Model(dbData).Updates(&data).Error
//	return &data, err
//}

// 状态更新
func (service *BaseServiceImpl[D, T]) UpdateStatus(id D, field string, fieldValue int) (bool, int64) {
	var data T
	affected := global.KSD_DB.Model(data).Where("id = ?", id).Update(field, fieldValue).RowsAffected
	return affected > 0, affected
}

// 自增
func (service *BaseServiceImpl[D, T]) IncrById(id D, field string) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" >= 0", id).Update(field, gorm.Expr(field+" + 1")).RowsAffected
	return affected > 0, affected
}

// 指定步长自增
func (service *BaseServiceImpl[D, T]) IncrByIdNum(id D, field string, fieldValue int) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" >= 0", id).Update(field, gorm.Expr(field+" + ?", fieldValue)).RowsAffected
	return affected > 0, affected
}

// 自减
func (service *BaseServiceImpl[D, T]) DecrById(id D, field string) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" > 0", id).Update(field, gorm.Expr(field+" - 1")).RowsAffected
	return affected > 0, affected
}

// 指定步长自减
func (service *BaseServiceImpl[D, T]) DecrByIdNum(id D, field string, fieldValue int) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" > 0", id).Update(field, gorm.Expr(field+" - ?", fieldValue)).RowsAffected
	return affected > 0, affected
}

// 批量自增 + 1
func (service *BaseServiceImpl[D, T]) Incrs(ids []D, field string) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" >= 0", ids).
		Update(field, gorm.Expr(field+" + 1")).RowsAffected
	return affected > 0, affected
}

// 批量自减 + 1
func (service *BaseServiceImpl[D, T]) Decrs(ids []D, field string) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" > 0", ids).
		Update(field, gorm.Expr(field+" - 1")).RowsAffected
	return affected > 0, affected
}

// 批量自增 + num
func (service *BaseServiceImpl[D, T]) IncrsByNum(ids []D, field string, num int) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" >= 0", ids).
		Update(field, gorm.Expr(field+" + ?", num)).RowsAffected
	return affected > 0, affected
}

// 批量自减 - num
func (service *BaseServiceImpl[D, T]) DecrsByNum(ids []D, field string, num int) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" > 0", ids).
		Update(field, gorm.Expr(field+" - ?", num)).RowsAffected
	return affected > 0, affected
}
