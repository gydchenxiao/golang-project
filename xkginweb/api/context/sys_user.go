package context

type SysUserContext struct {
	Account  string `json:"account" `  // 用户登录名
	Password string `json:"password" ` // 密码加盐
	Username string `json:"username"`  // 用户昵称
	Avatar   string `json:"avatar"`    // 用户头像
	Phone    string `json:"phone"`     // 用户手机号
	Email    string `json:"email"`     // 用户邮箱
}
