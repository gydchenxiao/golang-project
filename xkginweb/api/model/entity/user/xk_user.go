package user

import (
	"time"
	"xkginweb/global"
)

type XkUser struct {
	global.GVA_MODEL
	Nickname  string    `gorm:"size:60;not null;default:'';comment:昵称" json:"nickname"`
	Account   string    `gorm:"size:30;not null;default:'';comment:账号" json:"account"`
	Age       uint8     `gorm:"size:200;not null;default:0;comment:年龄" json:"age"`
	Sex       uint8     `gorm:"size:200;not null;default:2;comment:性别 0女 1男 2 保密" json:"sex"`
	Address   string    `gorm:"size:100;not null;default:'';comment:地址" json:"address"`
	Telephone string    `gorm:"size:20;not null;default:'';comment:电话号码" json:"telephone"`
	Birthday  string    `gorm:"size:30;not null;default:'';comment:生日" json:"birthday"`
	Avatar    string    `gorm:"size:200;not null;default:'';comment:头像" json:"avatar"`
	Password  string    `gorm:"size:60;not null;default:'';comment:密码" json:"password"`
	Forbbiden uint8     `gorm:"size:1;not null;default:0;comment:0未禁止1被禁止" json:"forbbiden"`
	Sign      string    `gorm:"size:100;not null;default:'';comment:签名" json:"sign"`
	Vip       uint8     `gorm:"size:1;not null;default:0;comment:0是普通用户 1 月会员  2 年会员 3 svip" json:"vip"`
	UnionId   string    `gorm:"size:30;not null;default:'';comment:微信登录联合id" json:"unionId"`
	OpenId    string    `gorm:"size:30;not null;default:'';comment:微信登录唯一openid" json:"openId"`
	Viptime   time.Time `gorm:"comment:vip过期时间" json:"viptime"`
	Longitude string    `gorm:"size:30;not null;default:'';comment:经度" json:"longitude"`
	Latitude  string    `gorm:"size:30;not null;default:'';comment:纬度" json:"latitude"`
}

func (XkUser) TableName() string {
	return "xk_user"
}
