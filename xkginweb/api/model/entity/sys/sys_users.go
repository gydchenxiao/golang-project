package sys

import "xkginweb/global"

// structs 属性的设置是为了：gorm框架更新多列中，更新0值失败的问题（因为修改了 is_deleted 软删除字段）
type SysUser struct {
	//ID        uint                  `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
	//CreatedAt time.Time             `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
	//UpdatedAt time.Time             `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
	//IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`
	global.GVA_MODEL `structs:"-"`
	UUID             string `json:"uuid" structs:"-" gorm:"index;comment:用户UUID"` // 用户UUID
	Slat             string `json:"slat" structs:"-" gorm:"comment:密码加盐"`         // 用户登录密码
	Enable           int    `json:"enable" structs:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	Account          string `json:"account" structs:"account" gorm:"index;comment:用户登录名"`                                               // 用户登录名
	Password         string `json:"password" structs:"password" gorm:"comment:用户登录密码"`                                                  // 密码加盐
	Username         string `json:"username" structs:"username"  gorm:"default:系统用户;comment:用户昵称"`                                      // 用户昵称
	Avatar           string `json:"avatar" structs:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone            string `json:"phone"  structs:"phone" gorm:"comment:用户手机号"`                                                        // 用户手机号
	Email            string `json:"email"  structs:"email" gorm:"comment:用户邮箱"`
}

func (s *SysUser) TableName() string {
	return "sys_users"
}
