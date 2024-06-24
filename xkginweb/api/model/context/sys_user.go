package context

import (
	"github.com/gookit/validate"
)

type SysUserContext struct {
	Account  string `validate:"required" json:"account" `                     // 用户登录名
	Password string `validate:"required|minLen:6|maxLen:16" json:"password" ` // 密码加盐
	Username string `validate:"required" json:"username"`                     // 用户昵称
	Avatar   string `validate:"required|fullUrl" json:"avatar"`               // 用户头像
	Phone    string `validate:"required|cnMobile" json:"phone"`               // 用户手机号
	Email    string `validate:"required|email" json:"email"`                  // 用户邮箱
	//Enable    int                   `validate:"required" json:"enable"`
	//IsDeleted soft_delete.DeletedAt `validate:"required" json:"is_deleted"`
}

// Messages 您可以自定义验证器错误消息
func (f SysUserContext) Messages() map[string]string {
	return validate.MS{
		"required":        "{field}不能为空",
		"Password.minLen": "{field}不能少于6位",
		"Password.maxLen": "{field}最大不能超过16位",
		"Email.email":     "{field}输入不合法",
		"Avatar.fullUrl":  "{field}输入不合法",
		"Phone.cnMobile":  "{field}输入不合法",
	}
}

// Translates 你可以自定义字段翻译
func (f SysUserContext) Translates() map[string]string {
	return validate.MS{
		"Account":  "用户账号",
		"Password": "用户密码",
		"Username": "用户昵称",
		"Avatar":   "用户头像",
		"Phone":    "用户手机",
		"Email":    "用户Email",
	}
}
