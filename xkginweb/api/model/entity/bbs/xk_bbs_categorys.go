package bbs

import "xkginweb/global"

type BbsCategory struct {
	global.GVA_MODEL
	Title       string `gorm:"size:100;not null;default:'';comment:分类名称" form:"title" json:"title"`
	Description string `gorm:"size:400;not null;default:'';comment:描述" form:"description" json:"description"`
	Parent_id   uint   `gorm:"not null;default:0;comment:父ID" form:"parentId" json:"parentId"`
	Sorted      int    `gorm:"not null;default:1;comment:排序" form:"sorted" json:"sorted"`
	Status      int8   `gorm:"size:1;not null;default:1;comment:0 未发布 1 发布" form:"status" json:"status"`
}

// 方法
func (BbsCategory) TableName() string {
	return "xk_bbs_category"
}
