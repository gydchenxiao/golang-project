package sys

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// 注意下面 structs 对应的是数据库表中的字段的哦
type SysMenus struct {
	ID        uint      `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
	CreatedAt time.Time `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
	UpdatedAt time.Time `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
	// 注意下面的 json:"isDeleted 是 isDeleted 不是 is_deleted 的哦
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`

	ParentId uint   `json:"parentId" gorm:"comment:父菜单ID"`      // 父菜单ID
	Path     string `json:"path" gorm:"comment:路由path"`         // 路由path
	Title    string `json:"title" gorm:"comment:菜单名称"`          // 菜单名称
	Name     string `json:"name" gorm:"comment:路由name 用于国际化处理"` // 路由name 用于国际化处理
	Hidden   int    `json:"hidden" gorm:"comment:是否在列表隐藏"`      // 是否在列表隐藏
	//Component string `structs:"-" json:"component" gorm:"comment:对应前端文件路径"`  // 对应前端文件路径
	Sort int    `json:"sort" gorm:"comment:排序标记"`     // 排序标记
	Icon string `json:"icon" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	// 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
	Children []*SysMenus `gorm:"-" json:"children"`
	//TopObj   *SysMenus   `gorm:"-" json:"-"`
}

func (s *SysMenus) TableName() string {
	return "sys_menus"
}
