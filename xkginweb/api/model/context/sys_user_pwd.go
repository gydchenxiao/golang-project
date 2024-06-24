package context

import "github.com/gookit/validate"

// 修改自己密码的使用
type SysUserPwdContext struct {
	Password        string `validate:"required|minLen:6|maxLen:16" json:"password" `        // 密码
	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword" ` // 确认密码
}

// 管理员修改用户数据的载体
type UserPwdContext struct {
	UserId          uint   `validate:"required|gt:0" json:"userId"`                         // 修改那个用户的id
	Password        string `validate:"required|minLen:6|maxLen:16" json:"password" `        // 密码
	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword" ` // 确认密码
}

// Messages 您可以自定义验证器错误消息
func (f UserPwdContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
		"gt":       "{field}必须大于0",
		"minLen":   "{field}大于等于6位",
		"maxLen":   "{field}小于等于16位",
		//"UserId.required":          "用户ID不能为空",
		//"Password.required":        "密码不能为空",
		//"ConfirmPassword.required": "确认密码不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f UserPwdContext) Translates() map[string]string {
	return validate.MS{
		"UserId":          "用户IdXXXX",
		"Password":        "密码XXXX",
		"ConfirmPassword": "确认密码XXXX",
	}
}

// Messages 您可以自定义验证器错误消息
func (f SysUserPwdContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f SysUserPwdContext) Translates() map[string]string {
	return validate.MS{
		"Password":        "密码",
		"ConfirmPassword": "确认密码",
	}
}
