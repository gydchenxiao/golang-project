package model

import (
	"github.com/gookit/validate"
)

type XkBbs struct {
	Title string
	Code  string
	Email string
}

//
//type XkBbs struct {
//	Title string `validate:"required|minLen:7"`
//	Code  string `validate:"sodeValidator"`
//	Email string `validate:"email"`
//}
//
//// Custom+Validator 定义在结构体中的自定义验证器
//func (f XkBbs) CodeValidator(val string) bool {
//	return len(val) == 4
//}
//
//func (f XkBbs) SodeValidator(val string) bool {
//	return len(val) == 6
//}
//
//// Messages 您可以自定义验证器错误消息
//func (f XkBbs) Messages() map[string]string {
//	return validate.MS{
//		"required":           "注意! 该字段 {field} 是必填的",
//		"Title.required":     "请输入{field}",
//		"Title.minLen":       "{field}必须大于等于7",
//		"Email.email":        "{field}不合法xxx",
//		"Code.sodeValidator": "{field}必须是6位",
//	}
//}

// Translates 你可以自定义字段翻译
func (f XkBbs) Translates() map[string]string {
	return validate.MS{
		"Title": "课程标题",
		"Email": "课程Email",
		"Code":  "登录验证码",
	}
}
