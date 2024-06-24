package video

import (
	"xkginweb/global"
)

type XkVideoCategory struct {
	global.GVA_MODEL
	CategoryName string `json:"categoryName" gorm:"not null;default:'';comment:分类名称"`
	Description  string `json:"description" gorm:"not null;default:'';comment:分类描述"`
	ParentId     uint   `json:"parentId" gorm:"not null;default:0;comment:分类的主ID"`
	Status       int8   `json:"status" gorm:"not null;default:1;comment: 0 未发布 1 发布"`
	Sorted       int32  `json:"sorted" gorm:"not null;default:1;comment:排序"`
	// 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
	Children []*XkVideoCategory `gorm:"-" json:"children"`
	TopObj   *XkVideoCategory   `gorm:"-" json:"-"`
}

func (XkVideoCategory) TableName() string {
	return "xk_video_category"
}
