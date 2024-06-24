package vo

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type SysUsersVo struct {
	ID        uint                  `json:"id"` // 主键ID
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	IsDeleted soft_delete.DeletedAt `json:"isDeleted"`
	Enable    int                   `json:"enable"`
	Account   string                `json:"account"`  // 用户登录名
	Password  string                `json:"password"` // 密码加盐
	Username  string                `json:"username"` // 用户昵称
	Avatar    string                `json:"avatar"`   // 用户头像
	Phone     string                `json:"phone"`    // 用户手机号
	Email     string                `json:"email"`    // 用户邮箱
	RoleIds   string                `json:"roleIds"`  // 用户角色ids
}
