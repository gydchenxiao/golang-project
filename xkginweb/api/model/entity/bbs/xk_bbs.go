package bbs

import "xkginweb/global"

type XkBbs struct {
	global.GVA_MODEL
	Title        string `json:"title" gorm:"not null;index;comment:标题"`
	Img          string `json:"img"  gorm:"not null;default:'';comment:封面图"`
	Description  string `json:"description" gorm:"not null;default:'';comment:描述"`
	Content      string `json:"content" gorm:"not null;default:'';comment:文章内容--MD格式"`
	HtmlContent  string `json:"htmlContent" gorm:"not null;default:'';comment:文章内容--MD格式"`
	CategoryId   uint   `json:"categoryId" gorm:"not null;default:0;comment:文章分类ID"`
	CategoryName string `json:"categoryName" gorm:"not null;default:'';comment:文章分类名称"`
	ViewCount    int8   `json:"viewCount" gorm:"not null;default:0;comment:文章阅读数"`
	Comments     int8   `json:"comments" gorm:"not null;default:0;comment:评论数"`
	CommentsOpen int8   `json:"commentsOpen" gorm:"not null;default:1;comment:是否允许评论 0 不允许  1 允许"`
	Status       int8   `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
	UserId       uint   `json:"userId" gorm:"not null;comment:文章作者ID"`
	Username     string `json:"username"  gorm:"not null;default:'';comment:文章发布者用户名"`
	Avatar       string `json:"avatar"  gorm:"not null;default:'';comment:文章发布者头像"`
}

func (XkBbs) TableName() string {
	return "xk_bbs"
}
