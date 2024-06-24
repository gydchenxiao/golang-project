package user

import "xkginweb/global"

// 作者表
type XkUserAuthor struct {
	global.GVA_MODEL
	UserId          int    `gorm:"column:user_id;not null;comment:用户ID;" json:"userId"`
	AuthorName      string `gorm:"column:author_name;size:60;not null;default:'';comment:作者名称" json:"authorName"`
	IdCard          string `gorm:"column:id_card;size:30;not null;default:'';comment:身份证" json:"idCard"`
	IdCardCover     string `gorm:"column:id_card_cover;size:200;not null;default:'';comment:身份证正面" json:"idCardCover"`
	IdCardBackCover string `gorm:"column:id_card_back_cover;size:200;not null;default:'';comment:身份证反面" json:"idCardBackCover"`
	Address         string `gorm:"column:address;size:100;not null;default:'';comment:收货地址" json:"address"`
	Bank            string `gorm:"column:bank;size:30;not null;default:'';comment:银行卡号" json:"bank"`
	Status          uint8  `gorm:"column:status;size:1;not null;default:0;comment:0未审核 1已审核 2已拒绝" json:"status"`
	RefuseReason    string `gorm:"column:refuse_reason;size:200;not null;default:'';comment:拒绝原因" json:"refuseReason"`
}

func (XkUserAuthor) TableName() string {
	return "xk_user_author"
}
